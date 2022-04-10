import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import api, { GetResponse } from "./../services/api";

export function PageCar(){
    let { id } = useParams();
    var obj = GetResponse("/api/car/1", true)

    return (
        <h1>Car a {obj.name}</h1>
    );
}