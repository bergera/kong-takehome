const { expect } = require("chai");
const supertest = require("supertest");

const server = supertest("http://localhost:8080");

describe("GET /services", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server.get("/services").set("X-User-Id", "nobody").expect(401).end(done);
  });

  it("should return an empty result when there are no services for the user's org", function (done) {
    server
      .get("/services")
      .set("X-User-Id", "4")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("count");
        expect(res.body.count).to.equal(0);

        expect(res.body).to.have.property("services");
        expect(res.body.services).to.be.empty;

        done();
      });
  });

  it("should return results for org_id=1", function (done) {
    server
      .get("/services")
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("services");
        expect(res.body.services).not.to.be.empty;
        expect(res.body.count).to.equal(res.body.services.length);
        expect(res.body.services[0]).to.have.property("serviceId");
        expect(res.body.services[0]).to.have.property("orgId");
        expect(res.body.services[0]).to.have.property("versionCount");
        expect(res.body.services[0]).to.have.property("title");
        expect(res.body.services[0]).to.have.property("summary");
        done();
      });
  });
});

describe("GET /services/:serviceID", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server.get("/services").set("X-User-Id", "nobody").expect(401).end(done);
  });

  it("should return 404 Not Found when the service does not exist", function (done) {
    server.get("/services/99999").set("X-User-Id", "1").expect(404).end(done);
  });

  it("should return 404 Not Found when the service belongs to a different org", function (done) {
    server.get("/services/1").set("X-User-Id", "3").expect(404).end(done);
  });

  it("should return 200 OK when the service exists and belongs to user's org", function (done) {
    server
      .get("/services/1")
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("serviceId");
        expect(res.body).to.have.property("orgId");
        expect(res.body).to.have.property("versionCount");
        expect(res.body).to.have.property("title");
        expect(res.body).to.have.property("summary");
        expect(res.body).to.have.property("versions");
        expect(res.body.versions).not.to.be.empty;
        expect(res.body.versionCount).to.equal(res.body.versions.length);
        expect(res.body.versions[0]).to.have.property("serviceId");
        expect(res.body.versions[0]).to.have.property("versionId");
        expect(res.body.versions[0]).to.have.property("summary");
        done();
      });
  });
});

describe("GET /services/:serviceID/versions", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server.get("/services").set("X-User-Id", "nobody").expect(401).end(done);
  });

  it("should return 404 Not Found when the service does not exist", function (done) {
    server
      .get("/services/99999/versions")
      .set("X-User-Id", "1")
      .expect(404)
      .end(done);
  });

  it("should return 404 Not Found when the service belongs to a different org", function (done) {
    server
      .get("/services/1/versions")
      .set("X-User-Id", "3")
      .expect(404)
      .end(done);
  });

  it("should return 200 OK when the service exists and belongs to user's org", function (done) {
    server
      .get("/services/1/versions")
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("serviceId");
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("versions");
        expect(res.body.versions).not.to.be.empty;
        expect(res.body.count).to.equal(res.body.versions.length);
        expect(res.body.versions[0]).to.have.property("serviceId");
        expect(res.body.versions[0]).to.have.property("versionId");
        expect(res.body.versions[0]).to.have.property("summary");
        done();
      });
  });
});

describe("GET /services/:serviceID/versions/:versionID", function () {
  it("should return 401 Not Authorized when the X-User-Id header is missing", function (done) {
    server.get("/services").unset("X-User-Id").expect(401).end(done);
  });

  it("should return 401 Not Authorized when the X-User-Id is invalid", function (done) {
    server.get("/services").set("X-User-Id", "nobody").expect(401).end(done);
  });

  it("should return 404 Not Found when the service does not exist", function (done) {
    server
      .get("/services/99999/versions/1")
      .set("X-User-Id", "1")
      .expect(404)
      .end(done);
  });

  it("should return 404 Not Found when the version does not exist", function (done) {
    server
      .get("/services/1/versions/99999")
      .set("X-User-Id", "1")
      .expect(404)
      .end(done);
  });

  it("should return 404 Not Found when the service belongs to a different org", function (done) {
    server
      .get("/services/1/versions/1")
      .set("X-User-Id", "3")
      .expect(404)
      .end(done);
  });

  it("should return 200 OK when the service exists and belongs to user's org", function (done) {
    server
      .get("/services/1/versions/1")
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("serviceId");
        expect(res.body).to.have.property("versionId");
        expect(res.body).to.have.property("summary");
        done();
      });
  });
});
