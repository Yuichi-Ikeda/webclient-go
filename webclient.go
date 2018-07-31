import (
    "fmt"
    "net/http"
    "log"
    "os"
    "time"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
     pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
        address     = "localhost:50051"
        defaultName = "world"
)

var msg = "init"

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your url path is %s. \n", r.URL.Path)

	// call gRPC service
    msg = callgrpc()
    fmt.Fprintf(w, "Greeting from gRPC service : %s", msg)
}

func callgrpc() string {
        conn, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := pb.NewGreeterClient(conn)

        // Contact the server and print out its response.
        name := defaultName
        if len(os.Args) > 1 {
                name = os.Args[1]
        }
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Greeting: %s", r.Message)

        return r.Message
}
