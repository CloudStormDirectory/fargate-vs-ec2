import http.server


PORT = 8000


def fib(n):
    a, b = 0, 1
    for i in range(n):
        a, b = b, a + b
    return a


def rec_fib(n):
    if n < 2:
        return n
    return rec_fib(n - 1) + rec_fib(n - 2)


funcs = {
    "fib": fib,
    "rec_fib": rec_fib,
}


def noop():
    pass


class FibHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        print(self.protocol_version, flush=True)
        _none, func, number = self.path.split("/")
        message = funcs.get(func, noop)(int(number))

        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        self.wfile.write(bytes(str(message), "utf8"))



server_address = ("", PORT)
httpd = http.server.HTTPServer(server_address, FibHandler)
httpd.serve_forever()

