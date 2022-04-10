import axios from "axios";
import { useEffect, useState } from "react";

const api = axios.create({
  baseURL: "http://localhost:3001",
});

export function GetResponse(url: string, unique = false){
    var response: any = []
    const [obj = [], setObj] = useState();

    useEffect(() => {
        api
          .get(url)
          .then((response) => setObj(response.data))
          .catch((err) => { console.error("ops! ocorreu um erro" + err) })
    }, []);

    if(unique){
        obj.forEach(element => {
            response = element;
        });
    }else{
        response = obj
    }
    return response
}

export default api;