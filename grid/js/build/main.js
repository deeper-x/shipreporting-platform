"use strict";
function getYear(value) {
    return Number(value.split("-")[0]);
}
function getMonth(value) {
    return Number(value.split("-")[1]);
}
function getDay(value) {
    return Number(value.split("-")[2].substr(0, 2));
}
function getHour(value) {
    var hhmm = value.split("T")[1];
    return Number(hhmm.split(":")[0]);
}
function getMinute(value) {
    var hhmm = value.split("T")[1];
    return Number(hhmm.split(":")[1]);
}


module.exports = {
    'getYear': getYear,
    'getMonth': getMonth,
    'getDay': getDay,
    'getHour': getHour,
    'getMinute': getMinute,
}