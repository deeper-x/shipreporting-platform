const firstElem = 0;
const secondElem = 1;
const thirdElem = 2;

export const getYear = (value: string): number => {
    return Number(value.split("-")[firstElem]);
}

export const getMonth = (value: string): number => {
    return Number(value.split("-")[secondElem]);
}

export const getDay = (value: string): number => {
    return Number(value.split("-")[thirdElem].substr(0,2));
}

export const getHour = (value: string): number => {
    if (value.indexOf("T") != -1) {
        const hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    } else {
        const hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    }
        
}

export const getMinute = (value: string): number => {
    if (value.indexOf("T") != -1) {
        const hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    } else {
        const hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    }
}

