function getYear(value: string): number{
    return Number(value.split("-")[0]);
}

function getMonth(value: string): number{
    return Number(value.split("-")[1]);
}

function getDay(value: string): number{
    return Number(value.split("-")[2].substr(0,2));
}

function getHour(value: string): number{
    let hhmm = value.split("T")[1];
    return Number(hhmm.split(":")[0]);
}

function getMinute(value: string): number{
    let hhmm = value.split("T")[1];
    return Number(hhmm.split(":")[1]);
}


