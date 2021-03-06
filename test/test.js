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

  it("should return 200 OK with default limit and offset", function (done) {
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
        expect(res.body.count).to.equal(5);
        expect(res.body.limit).to.equal(5);
        expect(res.body.offset).to.equal(0);
        expect(res.body.services).not.to.be.empty;
        expect(res.body.count).to.equal(res.body.services.length);
        expect(res.body.services[0]).to.have.property("serviceId");
        expect(res.body.services[0].serviceId).to.equal("1");
        expect(res.body.services[0]).to.have.property("orgId");
        expect(res.body.services[0]).to.have.property("versionCount");
        expect(res.body.services[0]).to.have.property("title");
        expect(res.body.services[0]).to.have.property("summary");
        done();
      });
  });

  it("should return page 1", function (done) {
    const count = 65;
    const limit = 5;
    const offset = 0;
    server
      .get(`/services?limit=${limit}&offset=${offset}`)
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("limit");
        expect(res.body).to.have.property("offset");
        expect(res.body).to.have.property("services");
        expect(res.body.services).not.to.be.empty;
        expect(res.body.count).to.equal(limit);
        expect(res.body.limit).to.equal(limit);
        expect(res.body.offset).to.equal(offset);
        expect(res.body.services[0]).to.have.property("serviceId");
        expect(res.body.services[0].serviceId).to.equal("1");
        expect(res.body.services[0]).to.have.property("orgId");
        expect(res.body.services[0]).to.have.property("versionCount");
        expect(res.body.services[0]).to.have.property("title");
        expect(res.body.services[0]).to.have.property("summary");
        done();
      });
  });

  it("should return page 2", function (done) {
    const limit = 5;
    const offset = 5;
    server
      .get(`/services?limit=${limit}&offset=${offset}`)
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        console.log(res);
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("limit");
        expect(res.body).to.have.property("offset");
        expect(res.body).to.have.property("services");
        expect(res.body.services).not.to.be.empty;
        expect(res.body.count).to.equal(limit);
        expect(res.body.limit).to.equal(limit);
        expect(res.body.offset).to.equal(offset);
        expect(res.body.services[0]).to.have.property("serviceId");
        expect(res.body.services[0].serviceId).to.equal("6");
        expect(res.body.services[0]).to.have.property("orgId");
        expect(res.body.services[0]).to.have.property("versionCount");
        expect(res.body.services[0]).to.have.property("title");
        expect(res.body.services[0]).to.have.property("summary");
        done();
      });
  });

  it("should match 'southern' and 'southwest' for query 'south'", function (done) {
    server
      .get(`/services?q=south`)
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("services");
        expect(res.body.services).not.to.be.empty;
        expect(res.body.count).to.equal(2);
        expect(res.body.services[0].serviceId).to.equal("24");
        expect(res.body.services[1].serviceId).to.equal("31");
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

  it("should return 200 OK with default limit and offset", function (done) {
    server
      .get("/services/6/versions")
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("serviceId");
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("versions");
        expect(res.body.count).to.equal(5);
        expect(res.body.limit).to.equal(5);
        expect(res.body.offset).to.equal(0);
        expect(res.body.versions).not.to.be.empty;
        expect(res.body.count).to.equal(res.body.versions.length);
        expect(res.body.versions[0]).to.have.property("serviceId");
        expect(res.body.versions[0]).to.have.property("versionId");
        expect(res.body.versions[0].versionId).to.equal("1");
        expect(res.body.versions[0]).to.have.property("summary");
        done();
      });
  });

  it("should return page 1", function (done) {
    const limit = 5;
    const offset = 0;
    server
      .get(`/services/6/versions?limit=${limit}&offset=${offset}`)
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("serviceId");
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("versions");
        expect(res.body.limit).to.equal(limit);
        expect(res.body.offset).to.equal(offset);
        expect(res.body.versions).not.to.be.empty;
        expect(res.body.count).to.equal(res.body.versions.length);
        expect(res.body.versions[0]).to.have.property("serviceId");
        expect(res.body.versions[0]).to.have.property("versionId");
        expect(res.body.versions[0].versionId).to.equal("1");
        expect(res.body.versions[0]).to.have.property("summary");
        done();
      });
  });

  it("should return page 2", function (done) {
    const limit = 5;
    const offset = 5;
    server
      .get(`/services/6/versions?limit=${limit}&offset=${offset}`)
      .set("X-User-Id", "1")
      .expect(200)
      .end(function (err, res) {
        if (err) {
          done(err);
        }
        expect(res.body).to.have.property("serviceId");
        expect(res.body).to.have.property("count");
        expect(res.body).to.have.property("versions");
        expect(res.body.limit).to.equal(limit);
        expect(res.body.offset).to.equal(offset);
        expect(res.body.versions).not.to.be.empty;
        expect(res.body.count).to.equal(res.body.versions.length);
        expect(res.body.versions[0]).to.have.property("serviceId");
        expect(res.body.versions[0]).to.have.property("versionId");
        expect(res.body.versions[0].versionId).to.equal("6");
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
