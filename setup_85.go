package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.5" )

func xcTFCgRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A64MdATHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OC1PH6ReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MfLGraGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T59ocikqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jKRM1vrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YjOHS2klWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 21jxsoD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OIMpUHAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hC7bmLsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YdGGcEPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IzISy4p9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxU7MKPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f5ihdb5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8TH4yYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SyFshzcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8QmT6mZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvNapesxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOGqUzH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EwUuv5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mI8EzHvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1M46cc4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YGZw7hAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D6A7IoAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uiNonP1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ON0q4ilrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntDWFNpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rr227RTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6VLu9sPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TiNT96RjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fo6QhLmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oi38qSBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9KsEERuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i9J4HihCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QPDYdV63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ATK5gzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHREvzh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JkerLOC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iASoQV2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VmYPlpVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQ9i9qZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OvBJWYhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sZmzaDwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBS6tA2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GAztAsBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OZcUawpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zhox5TYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8txubAJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4LwHxnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQtj2DGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ItKYkW4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0olYyBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6hx2BXGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54CHrwgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ME43XhezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8CJIKf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JcU85JFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rzUEOYCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fmdco7HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5v2fxERyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func djQGAnz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xYNrUF3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hk5i8xAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxlveBjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4yc5qj8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsmg7c0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r4ArTliIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wEjiu6ZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2kpAin79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYNjkQskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucV71BjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nrB0dFBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g7UaHjxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RGLZ6m3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BpiFBbv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MCWejWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MeWUb2KIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ClvEhpvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugLsvmrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yt2xdSIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kApcSAwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i6qP2ANkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pikzv6SXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xfsiqfOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZao8wKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func at8jkJLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fh47CpBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mD9QVxElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DHsW4vinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNOkfpq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ryH1LMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHdhgXpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pKY6bpbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func blTQTE2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xxFpxqZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func osRFjmQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNfItLGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1nx1YwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yVfCjrniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0jRQeNSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PfsjacBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JZwStSlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q0JECj1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1OHJ2H5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jqNcqusLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xvlpwo0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sX8tZtxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qj95ATIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xa4LYg8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y1tjilmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SSXks9ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRvAVujeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qNxQflTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1hIonttoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbqljR2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orjsPq1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZrSxYQcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wNhakfyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jm3ofGLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eZnS7zdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nL60nQDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vSEwv04YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyCRgCh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1vbBsrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tysgo5bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UymTpW1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W3IxmWlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k4aqfIv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJ5Ci5Y3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iF0TkROZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtKoN8HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V9sEmAMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41Xf9NMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8iMLOCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xlGuHshrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uw5zC4kaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MoNwzN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aFm9gi89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lE2MrfF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wy8EqmtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func toYSPg4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mYx5wv67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fiFHIzSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m9L91pufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xV55vfBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vUh1GzCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gExK1CJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQZAEEHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jfA0cG8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rFp97k2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FgAUf0SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MjZ9Xct9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iNunzUIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoHKr4luWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ac1sL6eOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYNSGRqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K6ytQxdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B60F7uxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bSh2xc5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wahZl89CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S4yhhQSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NinBcNOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7XMariI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LGKh2VIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M5sp7LVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rW9UBFROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9ZQClgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gfEx472Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vd8mMqdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IIGHWE6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B4qKh7HKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jU5LGSvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1CbNhyP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ROcmb72hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func txetnNSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IB1HnKAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRCm729nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yD1bZDMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QoND46glWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIURF3krWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBKptKhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pOIMr3eeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJ7ECzoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IAVdUyL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rls4rAewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B1appexPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SPEpThWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kPbB78KsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYi6hYwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UX7IUMNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w7oXXEa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rlK1SYnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lt3Ryv8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7ulraVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDPnmjBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sKMFgEtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6J3sPr9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xAjBTeKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqmhKPVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NrFy3jjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DGQGcco5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OHVT1Gd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OjAqxTxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUxtkS9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VNscfiCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KSfCnx6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oJzdsI2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YkofTUdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JozOzTVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BzqJ2cgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GWv0soLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ele3mhN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdTGd98fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hl71bZxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PSglPEVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0St2nGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JrMr1K8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HOHyR7fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EKZ5ANvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0QaXXFILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jwVHWLXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8RG3PkrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcRLTLaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3mdVHAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IprElzltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PerMatyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6WC7dw2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EF6WqfnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5mEm54GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBJjwb1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QNKDWXKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eoi7WVR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0GrdAh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gw4EkONeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e97oec6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHA7fDa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXoLV4IZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKxutYgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ept3BJg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JQUPRfvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxRyGiXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DFzmfeCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lz4GSMaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LvZcsNqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z2rLd5isWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJ80Zu29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KwQz6LiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mkx99asYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z72xgWepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKQZNLOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gTBcZ1G1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GwiVd3aPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0EiBUcceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0AGAfyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NlO4xH6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iCuW4cRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YU7g2pH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IIEYeI2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSKMa5TDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XkbfPiwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qMP2zQluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GikRZPo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHOU43IwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFUEzoxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tdjmHhgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MU7MpDJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JL6TFnKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CvMJEvvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TkWOmOyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2KoPH7eLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 69fLlCOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x1nCGctbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NkKG6PbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kVh3Vz66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mJV6gM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pq5MZDsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Js90T7GDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KTe2M6twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EK08VIaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0SfuEUNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AqnAlsUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ID9IJMcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1QinBmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iv1vgwSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rc8v7vAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 77hRrxJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A320be9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pKRC8KNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ofra5qt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJQUTP1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NKYS08FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Anxr4RHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func US194v4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func laIL1Ns1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJaq0iJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k4Er8Su3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E3gznOtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dw5bTpFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JHgxuVI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJxGYgmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDp5vYfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPWasZnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7aFTz7ZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMqst1hdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hD62POX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lFuQ2Be4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzDjxHvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MyzsVgfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUyN92YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ev59s9U8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwxk8fAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zj16MdiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jq3GU3C4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ymdwWzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7Eh4KrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9N2dEcWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joTLKeVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1oKoHTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tydMaJY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acwERC4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xGRVIJs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BEs2QXmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WY8yMJvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSohzdDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T9boNKciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AmfYjp4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6rzwZEMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pkSyeJxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LgWNF2mgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KUVsMRdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ne35ZhLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlnSeBbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8c6NTPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vne57ILjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RxGrJQmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vD7NHvw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAQFo75XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xwsUOpTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCAwcvQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oytTvofJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7QWI27yNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ruVaN1djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r8h320iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b9elX3mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cb4DGClVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YYKqo470Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mHUo5BtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3c9gE0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func prCxOvRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pa4xjpBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bBEw6Z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypdbKCB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pm618321Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ffDtm5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THzajs5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzGe6b31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HcQ70icgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUnslRvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ixjm9j3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2f3dAtmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oj9Nzd2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65MEWzGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xniveYaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T7YNc5m3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i9MGxfb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w1YzJNnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qhTIfBSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2Zejde7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jqbq4gUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjWtqYh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZc5onfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3WDE2U0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9V4mm2ANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZtdARQ5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcAUMgNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FDbh25ntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6jOqHND1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSeoenBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CSzhu4i4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkKFxRAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kFiVqZHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrCAKqnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8CmqJDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9MJv0glrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F7iaA55NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jBD4vsBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ri4L2v1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGfipp9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9w0US1UQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kcENXEkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dLqG8PCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TSF767F7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27Eie3WmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFhWPs2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eTmJ1JhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MhnYUy3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqVLatldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xBXMjNk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2m745wIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gDpAh9QBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5plPGUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAdIvJkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Svn4WqVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxRWNEX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GWqso523Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fFqIeu9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWHM5xz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GoztHElEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iKGH2XFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m8y8ryNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IqNresaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRQ7znjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHJy1h7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jq7EneuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0dEOugDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1L4VmYXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bseXIppCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4Aig1PFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OIvUB8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func riBGHRfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HoDXpq2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ku7bSJrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIPUiUXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8O4sszJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TNGF33KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EuSTyP6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyR4cND2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYgOmh3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b8IT0P9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whsKAEG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOfQAGMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWAGC08jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LXzklaZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03Y2uR8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YK5b1q9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kv5hpNekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Si4XGeAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3yHAz7CyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XLCWPmBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X661Y1OGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOxrzjnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lhDSQoAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l6DkkdokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bRCSWWtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BeI08CYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1KUwtimYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XghmWLqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhuQSwPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sd1efjnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qd7HEh4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6FNq28wzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E8H2tc6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6MwC8RwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNugo2q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eNj5BHTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6aA8pXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TpXMnCx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFE4Be6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLzbzXsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8TVA09uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Dx0VPnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LUa1ixjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tYqYL6n6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ymyWb58LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IfSNIxMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MMQkGq9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JCvchgMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wchqzv83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XcRWLZq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func akevcC7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i06JE875Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2QV6H70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ImEimSxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QebDVoxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func esN9G9CPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pcMDNKGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func miSNoq5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MI3swpuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mfb65gUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzKgtv9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1uhVSXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y0WERT3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OGWacOdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fgaWVKdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pwe9PELJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HWF8PGvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YepjiZM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4BXVXM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LiM7RiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q4RzWqAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFTJPu78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vyvp0yFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxlFjcNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZo7jbiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kClkMu7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBEvYUv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zvmyHYuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xnrYf7fcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vKY2JEzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Ver1HFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8gfF5PDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sEMuLIxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 86ydi1iuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9LQM2xv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmeH4yPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCoN8lWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ib5gtHjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sdlm1HIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvUlWPXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cM2JQ4GbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bEgEqBZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bib9mfKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KlUeXVqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 557lToIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pv2PtoJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JWGAruXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tJ2yTnkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BepnfeiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2yUjr63OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZeT4bUhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLSnI2WQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q0Fe2mWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xsZvGhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Sts3ApjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTGElVQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a5J7nwqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PSIyFQFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMutLEVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eEaHvgJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gz8WQwUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G8hOXIe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81jivn16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0oTn7jCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W9Tgr0hmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Inu3h4oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKVhf7IeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdAIOeipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3pW4I8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPAAyl0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fEYcj5k2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15e1chV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y41Oo3VLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VZpgqIMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZI3zWD5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O4nUpm3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 21F1YHLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6JVWULlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20x8Cz4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EMHMjBlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7aYwJkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func am21TNtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7MFkqpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvo1XFtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D8CpyijRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q7WAEv6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6kw074KdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUhXaxbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81HC1f3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dI1mLbK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fC9CyfdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuwqevYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxFzevwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g0pldQmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQ3isgGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fJsOUHTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NcxI1gFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RyFPenLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcgDQuV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrtnXUq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wz6wa1YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ArF2pk03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XySmfW3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u2P8CGCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MewQm6ccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ovio3qq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6HOJFiKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8eq1Q76FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hOwncGIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wEElZFRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MqnQqfcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 30RRXO4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nrPk8MhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hqJUtup4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7RiTO4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8zKnmbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQ1WnKv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8i7XyqthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opf1beIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxXzlJqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ySvgbU8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQxm6RYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vLDH0YpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBrYDREmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zCGZVx9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jts7Tq1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hhjCVXTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSLrXMU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nqpngKbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQ0XUNeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GhaQ2CzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRw87ScZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gW5zgJZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2G7urF0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ABewagBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aDsbPT4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CpA2VsBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xFUxqqpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCJ6LbsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Yk2IafCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEr5xfhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgBgi4WfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPuJfuyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eS8EymA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WKIwIJeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Je9nOPKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 53tbBg7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgsKoIaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rJN07cqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kYb2euCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tciHQViRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHUY7qQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yW6JHxM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhzgpCC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SxyU9IaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xDorUfZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kw6IW2xBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qadzhojAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Wofc3nbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A1u6eezIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A7hlMHYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fhdwmzrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HsBSKoHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RPXEJBREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCr3pLXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func onTHr5VOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHAhMDCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1dwWchqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qei2zCCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cRt0nGZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRRxQkRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8PWD4mksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPVfHaHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0S6Ahf1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ys9yjP7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxGp1cc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U75KXib2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YE5HIOQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IN09fna9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AUqn28cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8QkTVjqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func COZKWHpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5506cSAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B69eA6dBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tFdE7VN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7JCwDKmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ke8Gk1jJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QPAvww6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GeJwssYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sk9CHGvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDrIQUQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VeFiZHh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sSnGGnQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ENQKHkfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1CRSoztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObBxVxbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q5DbRRmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZvvXmKg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BflqK0CEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rPbAHn6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ohQmz2iAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lMaLHOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pFlT3enSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rlgr20VHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pNVhAteiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGgGNwrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLdrpqYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UY0RlUM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B15ut1VtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rg1AHqJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zSTMry79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PB3M5maYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bRX4dy1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7tvcZz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZRqTFm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nF7a0go1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUYp5ZgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGmKLOLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RfpZU1eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQlDPni8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1I3FwPQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7hfP8XJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iatdJzPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yP1QaINVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oaA2DdQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o6qWswohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func btaNovvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IAjI49pmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kz420TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15bsau8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nnE6E53FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X7jkossaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4MBuDubcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KWJopDkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FP8ALf2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SrYgt65kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZAZxpTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E0fYS4EPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KCiRqwPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rt6Cfw0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9IVyZlWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Axqt88vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func byVJoC4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxnQVt28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kr4LbeNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0E1lZ7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 904kFN8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6xSPn57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dx3sKLsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pdhAFxP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BbjodECrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E0Qfc4FRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voWG13rkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MAQnrFo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NKSzrzgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZU8X7ZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKpE21TLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZR14eqbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFdIpsxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wv6zRoWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9S4EE0WAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ST3YEVjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gcSkRMGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QNhB3TLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gb1bhiAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JHzywayuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IIT8FSIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tN8PHm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8t98FAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vOn0xuUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ef0yYUZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToOgyv7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdI8QV59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HEKU6s9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Poq1TMPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QHjlvBT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHJEEgDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bje0pWReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CuXvqzJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ayk5mEmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BekmWwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2CKwq1lZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5q2wz2SsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhAr0Q2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YGcSgN2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9uZfKIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHm17hk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6aucQXe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbXefRlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5c1wHuC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrZP6MC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJ9wpb7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2IiTPEgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wnogM6ISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KNpS2B7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jm0BuPhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y02LvEUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UE9pKq7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukMtwDTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zhVL0iukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qgXXxTtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVT8jNIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YcwT3iiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FVgUloCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1F027u7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GIREO856Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2heyBBpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QG44CeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvRcygF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPZTTbboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p577PR1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnrJZjALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wz3MoJHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MplXJROAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O0r8XjiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I8nDmSLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CkIkpq0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDur2wTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N5s7qh6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mx9xrk8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jyVztvaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JG4pBg3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u60zP8UPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rlnLOsZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MD9U0KYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSYfOCByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HjQvjSm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6WWBbMmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UF7TiG7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQPIXvy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VfH6jOHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcYwsEfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b0rLQOX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSB1inM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E0V5uu2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVKUHcfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8W3uD04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E751yYUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VXB1lUK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UDnYhUzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Mn350dVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jmMikzdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBnHeHTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AkFEKKKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func znSRatJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAaHrTDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGlV63ElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b0VvVbvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v46fy2hcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jz6rSBzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 70ZrBNzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xMhqQg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y16B7I3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i9gNkOxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dTNN8yZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VmTpueP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qji4zFnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1tUPE0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xr8IndK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dhjcZJ80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1qapef1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caD2JEc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jU6GB5kiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UO60hFUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yFMQsgfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTz1sB3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SDxWOV1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ISdkIAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWwJxDjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JY8Ok4EMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0i5CGf8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHkKig9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2qFkrpD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XytAlz6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VP31k5aiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZexEdv2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1fTjQFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yS7oU5gyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UeuKNO7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FPCPJakGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RYidJ47JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPQFev0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VltmJvG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mSqrGEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 96wTlnKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWTydmVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FPVRvcJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVzzVjlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSxiSmcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CU3hAsd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2S1swpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1QI2rwZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2mKl4m8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNpUKJfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SrqtmmLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOe9uQonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65wWYaSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s04YiDuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJIpxhGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e393B25HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BrFvFnu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NMywTBkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eV2Xmr95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AYLRBxRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wgPj7AeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cd28cUkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wc5mc45pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0zAGzk5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lNONIFyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ZwvCQvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRgEruIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NXbOlH9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JF8bbKPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pijIYv8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rDUBfOVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLakYsZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtipK1CmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WrESvCv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OazcIxQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 607EJUW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5G68LdzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D6jnkl8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1pDx9pEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YcSjOYQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvorRlBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IceCRzUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBtqlK9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fV7YvwkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lz7a7jEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxAN11o8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iUsb3RblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxBfNLCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func euZgN92OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func etcyUmexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBLBDvn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 10L5SM7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lBitgergWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MIrKzk89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nZOziO5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zgqFHbocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SjrtIaOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgxTKSY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K22qdGM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDcHBxGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x3LwXxoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sBCm89g3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OKPFqGhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MC49TJqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hax71bmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EbQ45Ym3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PsJarCZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s5aT5UwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func upYIMC2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IN5EbjrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avBdX40RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9MFey3qKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PuRqGNlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5NsfeuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cYMZSplSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cIgh1VS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBfFJVfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xubA2elXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RW5wVEpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4TvzFsbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UT3mZEnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z58fqEeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r8V5vK6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NEZYkluEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yjmaTT9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2qRQBEEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fW1a28t3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xP7tkEjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvADn6yfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2xSHB35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func juWvNZzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMiLzRjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPG4JBM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ETa5eV1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6X7fIvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zr2yyIbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpqRM1YgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func symvTnW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSe1grKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nlWymRSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IyiXGMBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Xf59Um6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T137nIQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pz20BPvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WXDpTBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZMsCkedaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFU14p9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0Gjlg8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cg0gW3ljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r003qhIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qmRKEdLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNagaSoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R7dxSeOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y7YvFfxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kxpiHg3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wE8FAcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LiyKfN0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SJEWR9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6GHWqB7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHkDLj5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5GHJhn6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uT1KsfUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThasIAmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7GCb8KfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CdpKvzSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IuDpPveNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4JYZHtK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ka1y8Gl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHS5YPj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEPzAHs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zbLXNN6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OQ4EJ67OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vq9Nc1OHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 69pUqjtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func meOu36XwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVGadSagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HO7qychfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0iA4fULSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bxqVMEwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IoifDZdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uPNFrtjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NaHxFVMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lZJm7qWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7KJDjYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSfFOe5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38juWCUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rvOpEXMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NcAXGiCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aDpmdMTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WpsE64yYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKSpeCGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4YLfqYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yQeggG16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pv9fFDoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdx64MafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2roKh0fbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0cEBgRdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LmFV987bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wLlGD0YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sCBlmFQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDGFae2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0fRE9fnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zk9gCbrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gcwkbnb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUVVZJThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFCnYn12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func seEsz216Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3DAdsCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qFb3Jle3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPef5bjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShthQ94NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8mN1tJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rV6qWY3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ntm6GuhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wr2wDaBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eTHHcgO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9hdYM3ioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3PhZ1I8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HeWbLtz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3gzNKu6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJMCwYbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TmzNMLqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gZX42ZuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VkzQ3DZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f20BnhdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MxiGJdP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJeonNZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vytCE2CcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ggiCKp6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y77Zu0SmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func duixqsdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2A3jROwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gRwTiunDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZtlOGtt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gnu0WVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20KzYhVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vc2kGtluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65ystAN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8briE6VTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMggJOvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgGSHTibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OkrQ7kyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iUXK2dsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqsI84ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O59GXpfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9EhS58SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tc7rxOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPZvkL3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjEZd6ATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwTVfS6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6I55extwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7c08wU8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sfq7C4BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5bhxcpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ZycDRzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GWiUtrTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k17neQ8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NPqvfA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7WT7QfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgY0CB0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOVycJJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6LLiNCG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gGlbNJ6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 94ID3JmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Vi55HFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3922RCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vRXwChyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 95CAtGA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQSU9dMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xxY6LL1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9hkghq4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZxCUPVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lH00hKBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lq6bhdkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRU1QS91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUYIFVZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoX9FmLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AOcLhow4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tcpbeigpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ylJsFHinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qe0SvZsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zw9KmvkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mVux4aFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6AitNPbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpl350DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31g7S4p8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AXi4ZguLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSZTN3tfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4D9IOesUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zhfzIdmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JYYZXc3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IzU23y0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPGXVhVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lzohrM39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E24iMnjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sC1mtYmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fuuzNFF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q7URqjuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TR3KFuzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dZN5W7GcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nZ2mbbX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0z3TUxpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J97iSLfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OZmba2VIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5XFIVFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hVZWD49dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWQor9ZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Es6k5vPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8t3NYa5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtXJOj8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cRJLAz5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e8OoE7waWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NgZP0eQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxiDsUfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MsI44XVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xg16ItiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSMbCWeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vkcXz6IKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dVX8kz8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cRogqzaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2vZ2ixKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EAO0uYmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8oXhQDVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LLHc0A5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJaJrtqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hdIe7IEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfSsj86uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dfer7I4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S3FO69N8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LJKkd3YtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LeC9TYoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hSaSMLdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2SJhBD6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t5nA8wJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4EYyzVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMcUnrMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DDA2cYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ymzU0uGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WJijGg77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Al8YyKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spPWRWAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8yu8UIbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oAqTTev0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1PTtBfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncFLpK65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v140H4lcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iowAfYl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kU3oVh9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aeVTf9qAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GN7CvX8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYYIXaYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvYD5IS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VGmbH0jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WdNPFSSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PTSyuRThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvofDnxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NiKyTnrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DsikWb4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rPhHoIZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3lfv76brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kpWSUqnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fOK6VUINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PebRyGOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7wqgTHtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6mb9mNITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLuj4JRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SYcRdeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w1Jc88vkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xp2sW6m8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lNfPjmcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPFImbFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s51htTliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVvzOMkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNt8WCZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UXexHgTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jzMkUblfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vaqoNhioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1FZyLGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dhn9f66pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCHfngUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wqc34n8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mki3QnROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ohA0URklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iMFb9PTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uSOqptEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RxosPEPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXbn4BxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQJI1GtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E9mCSNikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1TNsbQk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEIsNNEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Hddm0GHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6qKXqBBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OMV45nu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvnbZiKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gWmVH5FHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4lMiWA5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5a7rC2RJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gay48tTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l5zoINjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ixIWj3VOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BvbTo5ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9n3KsJ2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gyiMVDm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HauCpK5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1r5WDkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99G1pCaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SmTQli1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func otgNWgtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XB6oefXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e5V8R4D7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wp8aVvHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g9g0KZJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L2T1Pr8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IKyalvgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SgY6tbaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CQJ2T2x3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n9sj8QZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eou0XFptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rf3RNy2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSxUALJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LfhTEXHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cNDOefQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJKOiGoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neZrRUojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pr16SsLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kueuhquoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X77bfzVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ei6cu4aEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xleCCpk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GnuUS2uuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func waFDizyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sTh8CHFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LAtdUtTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBtHhThEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIOXwUeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8cVfCbv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jupqZeEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDrMtGlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F6hBYO4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aFJIF0MrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t723isGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSMnNeRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVW31J7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func is0vGqtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r8VAr2AGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U2Prsh0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LcEBslxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vA6faYwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RREMmthOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6iarVtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEIFDJQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGPAdeAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9afI5aqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pA7zscfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7CkoWmWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anYXIU9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uNV1veNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZmekAw45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g6pYcmwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0q9ZnJRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UIObXpTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pVNkN0VFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c4sAc4oBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0YgnXFwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kFcmMX9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xv4WRj1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lrBgV09uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pmcer4ulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BjHnQDc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0ls8MGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVuXTd6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HFO8fwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGA8iXVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DUiqJwLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kSCNxMTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zdVKTIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DjTG3JmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unXUAcunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qs8npCRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8cObOi5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZHvfyQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICjFln6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpL2S6cKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u81k0sxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GAwtZFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uqYW6uwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6tnNmjjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H2mLR1IvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 28cpHQ88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nP6wyMmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d7GerczFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1IEHEZhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSoOH6k3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nhRIqB80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bcAIQUbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KpaTMuZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYarQfPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V5tKnkesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNB8QNMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XoFOP6LKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XIwmzbACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hhrgdL6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RXbSF1S9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ds4tB0w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n9CfSXQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5d3bbiL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5e3rWtVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OrFRmTFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3pkQov3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZjY0DnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iC4EFO1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzWbRswqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bB3Zh8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 79JUuc9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWJ1T63UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fw9vjUcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T5su7ZGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DNrvit7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m2uLZce3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acLAZcL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dTe0ntwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func irH80kjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnZlcZSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1eVcyKwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gx4DkapHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zxq0beLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6qqIU4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4Tquz5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2B7Rh1ZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kM4UQcWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3M9AjIMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGBsGASWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crlXHiQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IxvrukLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZDh1djtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LKk4lnQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41wN075tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCs9dMfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pj7VPlOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EiqPY7OWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gatd2z3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func utQhhdkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qDV6btNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkxxRStTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SeoSks2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rWfcEc5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N5hxeTFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FDeeP9FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PF9knKj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func urd32EQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D40xF76pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vhmrs0fqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQGqLMD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVlZs2obWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sr3RQuxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PTUTmvfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmnFWv9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hvu0TvPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lo6NvmZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Q1vUL6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b3wxIcfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQqIyrhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6EMkJK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oeY8B2rbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wikIZnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gA30emOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjFSEYLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bGTTrA2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47B5fkI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 10qzrVRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tlPAVleXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kJhxGbTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mMRz0XekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mpnMbSeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f12BmiWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BWCtQXsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYx6tf1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DSHZMspxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vt5eQuWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6qMbkyGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9KND0UtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wnf0BqTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDnYph6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvB7EMqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12a7eTDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCbY1D0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AuHzGseCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CKlT6SyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZFg3F6nMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NG7lbKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XBYeePkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXDMzEQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b0yHsDf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ULKSVfcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gqoIifCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uk7dKrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7ft9nPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XVrszYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r4KVvaXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KucWeZqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yOhX3lfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxtpQDT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hgL2FYBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 254Hg86EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oLJ4rWKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wevZeoM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pbwWkrgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yTG93TJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PPRYbKcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vL9w9q5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZViSlUUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbPHEPx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qQQIaS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oc8LCPW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fYvNt1otWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i9ipaKauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PFf520ORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uUbwfMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f4WrA76uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V09xLP3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4FQLdpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nXfKT92UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehwfdwzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZWG1oi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEkpuKdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q4PVedC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3Q7pX9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJtvU1gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPC8zDOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ClFUaUpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdUzEiKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRBQBnn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OxbAwVQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGiHCEkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sgMjuvqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9q2HfgF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TUFAjdwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJC9u3E7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7IPwKNMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aw12cBfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qJpc9cXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2eSmIdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QDMHlAq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rqvyvaxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9z54DndCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s8BzGLp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AkvPZ0ptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KrTG9FgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEiPn2J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xPBGFQVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3e8quCrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pEW7Vr93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CxVegKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func woCMC8L7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bC5X4OmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sA9XiQY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxTadiBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9aaqaNEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6JPBgPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlxcB5kTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzZvyOHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WtzFdn6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZACzW2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIYxuWbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezV06rpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2vXmYIStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9RMTguzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HexrNoGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GgxCZPZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TC3Qmjz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JsVSbsioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEe0P3KQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B8UdFj7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vDE5XreWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsshcaJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVEAwPiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0YlhaYPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZm0S8EmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOSUBq9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLPfTRAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fZ4667kiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FMHrNGxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzxEfA7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tBoKPqDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y12Aeh5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HsTS5UIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yREn0VqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tLkTfmlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aK8Tpb3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p8fopre1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7HVD2exFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1Co6ywyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nD5IhufAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9hTwbaYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHzvnxxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QrhWozc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BrQpzjvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eEk0l6ToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSrXdW7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ggq6UblVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LocWcPpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKVOi8o5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fHNKFkceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZC7qBctqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvDzMhoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tAuvpPoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5h84xk2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIq0ikG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkeqSLxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvgiJOxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5cMYY77yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QKe3wCmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WhvR0S7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kimEuSh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGM8jrp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BGGeMSqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oB01zDPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9Ewa8pdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GIukn5odWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FQhr6xBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYbYWzFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZnJMyr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uTaMwSjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pLAjQROzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PogJ6vvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4mBKzU3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SbojDZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VsK5D8fCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TNKBLtzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Go30BTR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LqaF1tBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYa0T1yrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlEScuLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mv8QSAv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOTaQ43wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3bnUEmpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AcFjtfmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4mmoucUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l732wiZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5uqWv6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zeWAffPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Vzfxs3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3GWV1gbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func agtdqqcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bfoHIvlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGS8NjvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rsGExSZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3LY9gOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OY960CrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func poBQhc2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kQUhFafzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTrBUoMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dSqgOOjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2vgWSrIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLNMonIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8P5KWt03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygftNtzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wkEcB0GjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qftf0DrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZYIoMyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qXmivtmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ibpwpYveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qE6w7JNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0fsCcQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2quwRzoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RkLXtK37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rze41cyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBXpcBDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKDHiEEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hEKhsDMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXK1RA4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t7zK6AGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u94bfxOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMLVUbM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C6IKbT28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opFCKRiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1OjMFvUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LEUBBOHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NSMMFEZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Df8TF6zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JOHDeWGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b97uqPDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85lNs2zDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVlL0rZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G44GSgJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vX8R5RyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Kb8xPP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W3arnFlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FsbFd1dNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R6it2kmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j6ttgtjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nmXlcscYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQjE7ysuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wsUi0NL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3F6g4enGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbQwY0bDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a2fjr8YuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KUZEy518Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MPkLKNiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHYF7vNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LsilyBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CLWPH7riWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udK71rV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oAxFeOmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BkvBaVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQb7XZvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zSfvcaCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A2twwX3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SkSxTyIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJDAptq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rdNGbKoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0cNGuKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xq8JqCJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOkaC1uYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zcTNfMFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOuJnaGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLeWAckYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rw65A53LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jer5OO9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YcJHzyzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nHEB11xIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ClRtx9wIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R0NB8hQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEOBzYn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gko02XXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J9XJfXkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scP3mrjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUY8yufsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AAI6RrgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tGyFPQaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Psh7uqXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONXwnpPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S3NiV7t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PC7cBCzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HuUBJ32rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s4fCMSv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13veMSOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jy0K2HpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wqi0XgSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRddqJRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UkBpGviSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMobUm0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvtpSh4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kaWnuvngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cSjXi4jCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DJSauTTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twltzxw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8rjfywY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ufCw4nZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjZQmOCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dvUdfn66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WAYF75AcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dtngzPJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y5HkEUUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VP9uuC2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q6dAMwSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bvw1VcXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYJB9jt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uPmDDApUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 70jpDUELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zI1TmcBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rscM5KL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func feQqklxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r4yqYKUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXYWIqtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cFzBNd34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zb2ZEwjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orL0gcHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQ1AMnzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5UcEMeKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DWAe1WR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lAX1VSsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRq349WaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LK1S5uxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMuQuUtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXiClY0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTYWhGosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8h7JO8IcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4lGsxdX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03WND9CSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 79529pkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQlItTCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oN9lb5B1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cg6bvYXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ofJjMJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ko5iZV4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eSYwiFFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hyy6zFIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23LaVrBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ak4bmZxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ycJx6grNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tTUumLqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2NGhntnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fkOB9QUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jibO2ET1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gcRnMwddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uO0W0KtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A9CS6AMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func On0HLNghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOep9JEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HjCb1EYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQ2ys5QGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OQ5PvmYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uLFRo1XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func skDkkD10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3ba65U7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PVP2TDJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kr7ZVmFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RzJoJFnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4jjeAUOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDnyiLmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U30Xw8ilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWTDH4mEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1nIsfQEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gFi40UWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnsRX6zvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eeOXb9saWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5hlJ4yiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BsSKl4axWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U0zC75QtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YVxwiFF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z0jvXJxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oWU3ZY6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lw2xCv1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UdBKFjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPAWzCrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x4vqGqBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SKh3JyLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PRWHZYmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ER1H7gZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCDZSLITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8r5AQDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzAU1V1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NphaqPqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2sacVEt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v04VksyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GUhuFG8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ibO2ZTNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZuok7tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZwksQjaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DJ4e9K0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wp3Kq7hWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6IfRGtndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DqpVpeCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mJUll4EkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0XaANrJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4pRjIQi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJAgGz13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DpS4j8e9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UmrFijbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VUWkEMSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQq7XllFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IRcNJHQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5W9FzfqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5CRbvOq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yshC0lJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WxWgYEiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4meudJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axpbGrz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0xNSS03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzu3bJyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdjA0h4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AOIGlmQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44SjrTBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5IhLz9VUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jfAnWVFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mWDnpMXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQJQxTzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TV4PLmkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e10WGjN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XmX9ja8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kx0jT3x9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bUxLw67gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w2oPTCg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8rpfRKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYnFHUrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9gLS1uuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGBX38xPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xl2quUDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PhrxD2FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqtspDTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6YRCr3v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWxa0VucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 340tFLkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvIbG1oSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mVwW4zeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BoM53u0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2q3Et8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QfizYsTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSgkl9wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mLAKMATHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYATNDoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V88UeybHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOaQqEtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tYKvBYKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5HshBdUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func haNeRFTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5I04IxOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nO0xh7n7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3fXefT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7iyx031Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMdCJJnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cKdMCzFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c21k5m6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJcBY4IOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ueTO0L1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udI0cj1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S9JtBJJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5e5DeFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PvT0hqv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AXWSGpGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdhHvJ6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6InrdDM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RwfW7PmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HcjGetghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func poHmJtWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7rlXh6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YHhCnBXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oN8lE84fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7dMj8UNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8lmt1CdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nF9U41IPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s4ryrviAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RkEEEHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jWiwot2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpXjJh4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1OI56bbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BOiaHv6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lzYDR34uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fh1lVqRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dToZcba4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 88aUgOozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqJZEAFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nyDvSyvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cxBfU8NWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7GW9jg1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VqIUrMRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9XAsPawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtH9QwR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jY7iZkMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WS5IEBlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rpJg31H5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7pAdlzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func naS9vc8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aYfGSV2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 45JV5vbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fivinoR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MbMwgX08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nzGBzYnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TfTADFlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4pabB9keWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a9UTzwAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5pgTt1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kOU1Hp77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5DmBW6eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOXFAskJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bEgxCJmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tZdcdEYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZkxHBsXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICoEqKgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOmVB4CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X2jHVZfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sP5KgBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6l6ZAnV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q86n7SJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rg9MjEPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CqqWU8EYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xkAzYMu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltohPXH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GeARjjnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MxPylBjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCGqliL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2XHdXMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7wDJuvgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRo4qRE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qS0leFW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bvz2t6E8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NLPLUbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sc37GPVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZUxYZRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMOG8caLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BANIjUGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOIl0Oe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NonGwFYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nqi1IgiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBmNL7iNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twcowioBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hUXHADd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qRatxvokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7RBVDVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oDCG3FGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dMlJoVVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CdN0stgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MMfKGwm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BXsMSXelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yh6TyiTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hkb0egaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pakT5omeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func noyRmNYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1QHLz30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijDLNF2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gD2U1GgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZSz9WFjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNl0A8hxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1C63T83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQWhGZDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OEuIxWGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REXzIUaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUgQjkECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZGqjdwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHOu04VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ovOajKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgCspBFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sk9W3dpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EeV2bpAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0LJssAusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fSVL8IalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ROyA2R9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SpafpjFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsXYuJ7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dyScrrgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RXK2YfpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5fht6fJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6m9oAeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GwkBFIDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTmVSrs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPxNoTdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bfUt9kaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6Nuku1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d6MxXDHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BlZQDyUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUWgMpDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FLQFH6VYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2a66TqoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nqxpZMBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B71kSL5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uWtqq0ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cZp71PdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HjJQ7S0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGzoBHOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KbalLfRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hGNh1NnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vWVOesCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQk4F7TEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DREAAvThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bd3DYBAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfJFvoisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DBwXOCa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SYGCaoCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNdSE7ImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CFitEzIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJe2KeHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4bqWNpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J3HjfecxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0JHbVxyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9es9DF2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0CRvYlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I1Ig95BpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vE3jbdrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNyNiZuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cE7Bhf4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yw9OYDtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VtnmOTVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HZ2t9u2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0uoFkoVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DpS9fFYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bcCnHyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g5dTL6fKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3THqKl5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zsFDXXNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzYXKqODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xE4U1JJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mUKIt6j8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KcpljvqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s717BeiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UpMbqNplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dY9g6ResWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r9EeNljIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5gU4aaDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heleteCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVK6U6s5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNo9OKSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vg2rWsq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sczGRw3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmRG83UPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7dKW5VY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TtO2gdIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVqfSG6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YA6141YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSlPtQqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FwxXeAzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p8yaPlwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func om3wXzYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAA0EAiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3mFhCJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iqlsJ98eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mn2IdVBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zm9dYqi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y19GGB9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npyPXZcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OhOo86TbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xRGKwTjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4RrZ5vsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gyQ3wH3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfAgGqRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DqOYYyYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d20MBF7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8Rd7fVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qszKGAjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLOIJJC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X82fsiq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vr6gN68CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F2TmID0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fAmd9hS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqiXk6jDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vu235YiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujQh5vABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpq8qFMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aUPf4F2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1YhgaISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3GpRNCUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J3uNAbGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aKl5RFVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Utw2NGe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K3CPyXOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kP2qZ14PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WfemwYDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func af5BEmwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rm3Fh9ECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMXF75yGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNGl0tcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dudcb2j9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ouj3EjIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zQxdgJAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7Ficre9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHIvH4W6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXQRZXStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xA5MZsSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cH0ieBBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dc5N9erWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0Z2esLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h1BN99zqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9RZpklhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IH2cCPQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTbyUeNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSSzskz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFw9WRFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exWlLoXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ULXKidLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6bgqrgz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func luaYFMe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8xF9uqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1l5hFwEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVHXXVh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pn4escBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Htk2NZ8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHLW8TudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N95bN2ixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ksHbKPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOoxSkmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zfj9yHERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qk7DB2cTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o2RtJuQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vkOGcNtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BpjkJ024Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1Phi5iTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mbDavIenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HkvbaKr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lTbNPiBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBD8Iy5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5mRIPxASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZyFF3wzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voUibYO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a97vTixEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QRnrSJfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1jpmoRtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e9lf7MBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Usd2JnICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jK6D9K3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jaqR78JaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MtI6xVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6KBH5fz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJpWn1vvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pyDrWUwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G4rOVauYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pm1dK8TTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TqB2p1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZotYBcsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYi0mDOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84mGijNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ww67ddbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ypuXalKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5NvNGL5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LlchwCuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcXRzBorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxnPmdu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCQdqVQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nQOj6QwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OwhX5wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5fOSqmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TejtRLZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eUmbx19ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwneG6r2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T5FKnxszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXpcMPHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UAlaOSoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DG1jeJf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TtmDwFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yHk0726mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCNnhTbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M10C5ZBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EWmiFnQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PlUPd4YvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9TFfY7GQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DU4DQcyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5uPQn2sjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7ywhtGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47NW2tWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q6fLRt56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tIyzArqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zeoV3nv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKYEo3AGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9sCnI0JhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gDHq9W8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27zEYRsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VAvIRLUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4yLSp6x4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xTtcjAuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iQ8zHxu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rN8AMoKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqkwGRoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivsGlQ8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oh90BWNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RXvyR2gnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vF5Nl3lPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TIvAv0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func La1fC6DSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r2EHTgf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xnO7PapJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d7SlOLj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vEh6nrSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfNfJNkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDDnyUb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5MhZvqcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wEjpKjSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HfCHMTbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lbDlPNsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwjVv78yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQhBAXiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLg35MWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNKvyUL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D80hbKRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axb9NBLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y30VyyEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0CrT4CQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyBYIV1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23r5p9b7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypL6Sa2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpEZzCjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 78ifYLXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0KaewqUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbp62t8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijSfGqH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4OWEWNvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nikEazeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgG5zvMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FTDhlBiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WbEx4orjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EvYJ2ol2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7JbN8PlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yjg9ynikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6XRCMDLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdpsBRTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AIn4gsaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uc378cu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sg2wrruuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rWgN5mfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PEWYonJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VsM7IUziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lP6K51fwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNQrjtFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nc6LDhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gKzawL3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8dTRUIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GefMPqF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whG8nC2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 196FYReSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func buuyTSrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLkWPY5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3F3GNB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryTdtUjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y7o06YTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xImnQJtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UQ2SZIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MlrOsZ1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vUyg5nkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQypZQOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OAKCzT4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mbWJkTvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55ZOgg8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7cBGSRGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xFQva7jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REkmZe2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FHI54iGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRJ0SXZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lLIprc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xbGGipv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6h9NAPvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func icZkXkH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4DJ6HcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DesGkP7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TU8cmEsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jK2ka6aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MuFL6i8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oh3PUcbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Eb3wkJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBNZVkvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func slv5qVweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F2w0egChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1idKSjG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CSGMrUy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rt01IqCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mJuLf2gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sL1MxjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BhVOTWulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lz6u4S4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0G6r8pWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lln3s38qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sB341GsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OUyX17MhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQkKQSrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zu9ayHdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6hIVJSgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YhKN5eJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MEkyq2CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I2oNyqGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBu7kAOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzhyMoSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orlBfU6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y7WiMaTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0PQX9UPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ge7tGEnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rhwWaDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVQna5f4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rY1GunsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ligvPFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YYxn5E5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rBwusyQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJwim7dWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 21pkTMXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WllVavqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IHY3tCcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVFR6YR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HHA9t5n1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PX4ZMbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VNACdNEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CjFAeaIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocgo0UjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KSuUtZl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D2Zi7yaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nk93DPuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UGx76CDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqkYgkr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGAA3znEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g6RiCyHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SbZWMyD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PCIjx2MRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1XYmRIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJqqdaHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oithf291Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jblm2n28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1dSq1QOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmfGCO40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3dZAFEBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BeCkgfjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zJw2O5KzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OdU2gzDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cxqwu1vGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aeDk438cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 93r7UgpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tI2u5qGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXav92tVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yEesMOJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xGdgCL4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e8m4FdMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMcmdARYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lxkfCrxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fStuYORYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FdmKeZaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

