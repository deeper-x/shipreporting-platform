var chai = require('chai');
var moduleMain = require('../build/main.js');

var inputString = "2019-10-31T02:21";

describe('Timestamp substring', function() {
    it('should return 2019', function() {
      chai.expect(moduleMain.getYear(inputString)).to.equal(2019);
    });
});

describe('Timestamp substring', function() {
  it('should return 10', function() {
    chai.expect(moduleMain.getMonth(inputString)).to.equal(10);
  });
});

describe('Timestamp substring', function() {
  it('should return 31', function() {
    chai.expect(moduleMain.getDay(inputString)).to.equal(31);
  });
});

describe('Timestamp substring', function() {
  it('should return 2', function() {
    chai.expect(moduleMain.getHour(inputString)).to.equal(2);
  });
});

describe('Timestamp substring', function() {
  it('should return 21', function() {
    chai.expect(moduleMain.getMinute(inputString)).to.equal(21);
  });
});
