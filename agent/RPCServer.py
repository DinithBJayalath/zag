from grpcService.nlpAgent_pb2_grpc import NLAgentServicer, add_NLAgentServicer_to_server
from grpcService.nlpAgent_pb2 import LLMResponse
import grpc
from concurrent import futures
from main import processCommand

class NLAgentService(NLAgentServicer):
    def SendPrompt(self, request, context):
        response = processCommand(request.prompt)
        return LLMResponse(response = response)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_NLAgentServicer_to_server(NLAgentService(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Server stated")
    server.wait_for_termination()

if __name__ == "__main__":
    serve()