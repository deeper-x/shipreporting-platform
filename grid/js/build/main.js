export function getYear(value) {
    return Number(value.split("-")[0]);
}
export function getMonth(value) {
    return Number(value.split("-")[1]);
}
export function getDay(value) {
    return Number(value.split("-")[2].substr(0, 2));
}
export function getHour(value) {
    if (value.indexOf("T") != -1) {
        var hhmm = value.split("T")[1];
        return Number(hhmm.split(":")[0]);
    }
    else {
        var hhmm = value.split(" ")[1];
        return Number(hhmm.split(":")[0]);
    }
}
export function getMinute(value) {
    if (value.indexOf("T") != -1) {
        var hhmm = value.split("T")[1];
        return Number(hhmm.split(":")[1]);
    }
    else {
        var hhmm = value.split(" ")[1];
        return Number(hhmm.split(":")[1]);
    }
}
