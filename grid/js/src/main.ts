export function getYear(value: string): number{
    return Number(value.split("-")[0]);
}

export function getMonth(value: string): number{
    return Number(value.split("-")[1]);
}

export function getDay(value: string): number{
    return Number(value.split("-")[2].substr(0,2));
}

export function getHour(value: string): number{
    let hhmm = value.split("T")[1];
    return Number(hhmm.split(":")[0]);
}

export function getMinute(value: string): number{
    let hhmm = value.split("T")[1];
    return Number(hhmm.split(":")[1]);
}

