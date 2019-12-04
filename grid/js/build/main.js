var firstElem = 0;
var secondElem = 1;
var thirdElem = 2;
var notFound = -1;
export var getYear = function (value) {
    if (value) {
        return Number(value.split("-")[firstElem]);
    }
    else {
        return "";
    }
};
export var getMonth = function (value) {
    if (value) {
        return Number(value.split("-")[secondElem]);
    }
    else {
        return "";
    }
};
export var getDay = function (value) {
    if (value) {
        return Number(value.split("-")[thirdElem].substr(firstElem, thirdElem));
    }
    else {
        return "";
    }
};
export var getHour = function (value) {
    if (value.indexOf("T") != notFound) {
        var hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    }
    else if (!value) {
        return "";
    }
    else {
        var hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    }
};
export var getMinute = function (value) {
    if (value.indexOf("T") != notFound) {
        var hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    }
    else if (!value) {
        return "";
    }
    else {
        var hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    }
};
export var buildGrid = function (uri, gridOptions) {
    agGrid.simpleHttpRequest({ url: uri }).then(function (data) {
        gridOptions.api.setRowData(data);
    });
};
