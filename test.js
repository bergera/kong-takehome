const supertest = require("supertest");

const server = supertest("http://localhost:8080");

describe("GET /services", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").send().expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server
      .get("/services")
      .set("X-User-Id", "nobody")
      .send()
      .expect(401)
      .end(done);
  });

  it("should return 501 Not Implemented", function (done) {
    server.get("/services").set("X-User-Id", "1").send().expect(501).end(done);
  });
});

describe("GET /services/:serviceID", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").send().expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server
      .get("/services")
      .set("X-User-Id", "nobody")
      .send()
      .expect(401)
      .end(done);
  });

  it("should return 501 Not Implemented", function (done) {
    server
      .get("/services/1")
      .set("X-User-Id", "1")
      .send()
      .expect(501)
      .end(done);
  });
});

describe("GET /services/:serviceID/versions", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").send().expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server
      .get("/services")
      .set("X-User-Id", "nobody")
      .send()
      .expect(401)
      .end(done);
  });

  it("should return 501 Not Implemented", function (done) {
    server
      .get("/services/1/versions")
      .set("X-User-Id", "1")
      .send()
      .expect(501)
      .end(done);
  });
});

describe("GET /services/:serviceID/versions/:versionID", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").send().expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server
      .get("/services")
      .set("X-User-Id", "nobody")
      .send()
      .expect(401)
      .end(done);
  });

  it("should return 501 Not Implemented", function (done) {
    server
      .get("/services/1/versions/1")
      .set("X-User-Id", "1")
      .send()
      .expect(501)
      .end(done);
  });
});
