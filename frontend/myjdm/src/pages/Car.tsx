import axios from "axios";
import { useEffect } from "react";
import { useParams } from "react-router-dom";

export function PageCar(){
    let { id } = useParams();

    useEffect(() => {
        axios
            .get("http://192.168.1.106:3001/api/car/1")
            .then(res => console.log(res.data.Value))
            .catch(err => console.log(err))
        
    }, []);

    return (
        <h1>Car {}</h1>
    );
}