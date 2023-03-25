package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/soniah/gosnmp"
)

func main() {
    // Replace with the IP address or hostname of the SNMP agent
    target := "192.168.0.1"

    // Replace with the SNMP community string
    community := "public"

    // Create an SNMP client
    params := &gosnmp.GoSNMP{
        Target:        target,
        Community:     community,
        Timeout:       time.Duration(5) * time.Second,
        MaxOids:       gosnmp.MaxOids,
        Version:       gosnmp.Version2c,
        Retries:       0,
        ExponentialTimeout: true,
        NonRepeaters:  0,
        MaxRepetitions: uint32(gosnmp.MaxRepetitions),
    }

    err := params.Connect()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error connecting to %s: %s\n", target, err)
        os.Exit(1)
    }
    defer params.Conn.Close()

    // Replace with the OID to start the SNMP walk from
    startOID := "1.3.6.1.2.1"

    // Perform the SNMP walk
    var wg sync.WaitGroup
    var results = make(chan gosnmp.SnmpPDU, 100)
    oid := startOID
    for {
        wg.Add(1)
        go func(oid string) {
            defer wg.Done()
            walkResult, err := params.BulkWalk(oid)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error walking %s: %s\n", oid, err)
                return
            }
            for _, pdu := range walkResult {
                results <- pdu
            }
        }(oid)

        nextOID, err := getNextOID(oid)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error getting next OID after %s: %s\n", oid, err)
            break
        }
        if nextOID == "" {
            break
        }
        oid = nextOID
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    // Print the SNMP walk results
    for pdu := range results {
        fmt.Printf("%s = %s\n", pdu.Name, pdu.Value)
    }
}

func getNextOID(oid string) (string, error) {
    parts := strings.Split(oid, ".")
    lastPart := parts[len(parts)-1]
    num, err := strconv.Atoi(lastPart)
    if err != nil {
        return "", err
    }
    if num == 0 {
        return "", nil
    }
    num--
    parts[len(parts)-1] = strconv.Itoa(num)
    return strings.Join(parts, "."), nil
}
