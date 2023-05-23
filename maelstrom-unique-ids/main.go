package main

import (
    crand "crypto/rand"
    mrand "math/rand"
    "encoding/json"
    "log"
    "math"
    "math/big"
    "strconv"
    "time"

    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
    node := maelstrom.NewNode()

    // handler callback function for "generate" msg
    node.Handle("generate", func(msg maelstrom.Message) error {
    // Unmarshal the message body as an loosely-typed map.
    var body map[string]any
    if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
    }

    // Update the message type to return back.
    body["type"] = "generate_ok"

    seed := mrand.NewSource(time.Now().UnixNano())
    gen := mrand.New(seed)
    n := gen.Intn(math.MaxInt64)
    
    // just try to introduce more randomness
    s := ""
    for i := 1; i <= 100; i++ {
        s += string(rune(gen.Intn(math.MaxInt64) % i))
    }

    // gather randomness
    randomID := strconv.Itoa(n) + s + strconv.Itoa(int(nrand()))
    body["id"] = randomID

    // Echo the original message back with the updated message type.
    return node.Reply(msg, body)
    })

    if err := node.Run(); err != nil {
        log.Fatal(err)
    }
}

// from mit labs
func nrand() int64 {
    max := big.NewInt(int64(1) << 62)
    bigx, _ := crand.Int(crand.Reader, max)
    x := bigx.Int64()
    return x
}

