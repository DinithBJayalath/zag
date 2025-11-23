from nlpAgent_pb2_grpc import NLAgentServicer, add_RouteGuideServicer_to_server
from nlpAgent_pb2 import LLMResponse
import grpc
from concurrent import futures
from main import processCommand

class NLAgentService(NLAgentServicer):
    def sendPrompt(self, request, context):
        response = processCommand(request.prompt)
        return LLMResponse(response = response)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_RouteGuideServicer_to_server(NLAgentService, server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()