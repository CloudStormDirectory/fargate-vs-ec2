const url = require("node:url");
const { createServer } = require("node:http");

const PORT = 8080;

function fibonacci(n) {
  let a = 0;
  let b = 1;

  for (let i = 0; i < n; i++) {
    a = b;
    b = a + b;
  }

  return a;
}

function recursiveFibonacci(n) {
  if (n < 2) return n;
  return recursiveFibonacci(n - 1) + recursiveFibonacci(n - 2);
}

const functions = {
  fib: fibonacci,
  rec_fib: recursiveFibonacci,
};

function fibHandler(request, response) {
  const parsedUrl = url.parse(request.url, true);
  const path = parsedUrl.pathname;

  if (request.method === "GET") {
    const [_, func, number] = path.split("/");
    const message = functions[func](number);

    response.writeHead(200, { "Content-Type": "text/plain" });
    response.end(String(message));
    return;
  }
  response.writeHead(404, { "Content-Type": "text/plain" });
  response.end(`Not found ${request.method} in ${path}`);
}

createServer(fibHandler).listen(PORT, () =>
  console.log(`listening on port ${PORT}`)
);
