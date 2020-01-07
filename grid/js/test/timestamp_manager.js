/* eslint-disable no-undef */
import { getYear, getMonth, getDay, getHour, getMinute } from '../src/main.js'
import { expect } from 'chai'

var inputString = '2019-10-31T02:21'

describe('Timestamp substring', function () {
  it('should return 2019', function () {
    expect(getYear(inputString)).to.equal(2019)
  })
})

describe('Timestamp substring', function () {
  it('should return 10', function () {
    expect(getMonth(inputString)).to.equal(10)
  })
})

describe('Timestamp substring', function () {
  it('should return 31', function () {
    expect(getDay(inputString)).to.equal(31)
  })
})

describe('Timestamp substring', function () {
  it('should return 2', function () {
    expect(getHour(inputString)).to.equal(2)
  })
})

describe('Timestamp substring', function () {
  it('should return 21', function () {
    expect(getMinute(inputString)).to.equal(21)
  })
})
