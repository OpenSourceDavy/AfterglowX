import React, {useEffect, useState} from "react";
import "leaflet/dist/leaflet.css";
import L from "leaflet";
import "leaflet.heat";
import CssBaseline from '@mui/material/CssBaseline';

import Container from '@mui/material/Container';
import Card from '@mui/material/Card';

import CardContent from '@mui/material/CardContent';
export default function Leafheatmap({userLat,userLng}) {


    // const addressPoints = [
    //     [43.653225, -79.383186, "571"],
    //
    // ]
    // setCordinates([props.userLat
    //     , props.userLng])
    const setpoints=(userLat,userLng)=>{
        let points=[]
        let i =0
        while(i<0.3){
            points.push([userLat+(Math.random(-1,1)),userLng+(Math.random(-1,1)),Math.random()])
            points.push([userLat-(Math.random(-1,1)),userLng+(Math.random(-1,1)),Math.random()])
            points.push([userLat+(Math.random(-1,1)),userLng-(Math.random(-1,1)),Math.random()])
            points.push([userLat-(Math.random(-1,1)),userLng-(Math.random(-1,1)),Math.random()])
            i+=0.001
        }
        return points;
    }
    let points = setpoints(userLat,userLng)

    useEffect(() => {
        let container = L.DomUtil.get('map');
        if(container != null){
            container._leaflet_id = null;
        }
        let map = L.map("map",{ scrollWheelZoom: false ,zoomControl: false}).setView([userLat,userLng], 11);

        L.tileLayer("https://{s}.tile.openstreetmap.fr/hot/{z}/{x}/{y}.png", {
            attribution:
                '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);

        // const points = addressPoints
        //     ? addressPoints.map((p) => {
        //         return [p[0], p[1]];
        //     })
        //     : [];


        // L.heatLayer([
        //     [userLat,userLng, 0.6], // lat, lng, intensity
        //     [43.653225, -80.383186, 0.5],
        //
        // ], {radius: 25, maxZoom:10}).addTo(map);


        L.heatLayer(points, {radius: 70, maxZoom:5, gradient:{0.1: 'yellow', 0.15: 'white',0.23:'grey' },blur:50}).addTo(map)


    });
    // return <div id="map" style={{ height: "100vh" }}></div>;
    return (

            <Container maxWidth= 'xl' sx={{ width: "1000vh" }}>
                <div id="map" style={{ height: "100vh" }}></div>;
            </Container>

    );
}