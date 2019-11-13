var firstElem = 0;
var secondElem = 1;
var thirdElem = 2;
export var getYear = function (value) {
    return Number(value.split("-")[firstElem]);
};
export var getMonth = function (value) {
    return Number(value.split("-")[secondElem]);
};
export var getDay = function (value) {
    return Number(value.split("-")[thirdElem].substr(0, 2));
};
export var getHour = function (value) {
    if (value.indexOf("T") != -1) {
        var hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    }
    else {
        var hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    }
};
export var getMinute = function (value) {
    if (value.indexOf("T") != -1) {
        var hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    }
    else {
        var hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    }
};
