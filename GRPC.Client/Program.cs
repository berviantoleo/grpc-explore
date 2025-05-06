using System.Threading.Tasks;
using Grpc.Net.Client;
using GRPC.Client;

// The port number must match the port of the gRPC server.
using var channel = GrpcChannel.ForAddress(Environment.GetEnvironmentVariable("SERVER_URL") ?? "http://localhost:5255");
var client = new Greeter.GreeterClient(channel);
var reply = await client.SayHelloAsync(
                  new HelloRequest { Name = "GreeterClient" });
Console.WriteLine("Greeting: " + reply.Message);
