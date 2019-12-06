const firstElem = 0;
const secondElem = 1;
const thirdElem = 2;
const notFound = -1;

export const getYear = (value: string): number|string => {
    if (value){
        return Number(value.split("-")[firstElem]);    
    } else {
        return "";
    }
}

export const getMonth = (value: string): number|string => {
    if (value){
        return Number(value.split("-")[secondElem]);
    } else {
        return "";
    }
        
}

export const getDay = (value: string): number|string => {
    if (value){
        return Number(value.split("-")[thirdElem].substr(firstElem,thirdElem));
    } else {
        return "";
    }
        
}

export const getHour = (value: string): number|string => {
    if (value.indexOf("T") != notFound) {
        const hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    } else if (!value) {
        return "";
    } else {
        const hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[firstElem]);
    }
        
}

export const getMinute = (value: string): number|string => {
    if (value.indexOf("T") != notFound) {
        const hhmm = value.split("T")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    } else if (!value) {
        return "";
    } else {
        const hhmm = value.split(" ")[secondElem];
        return Number(hhmm.split(":")[secondElem]);
    }
}


export const buildGrid = (uri: string, gridOptions: object) => {    
    agGrid.simpleHttpRequest({url: uri}).then(function(data:string) {
          gridOptions.api.setRowData(data);
    });
};
