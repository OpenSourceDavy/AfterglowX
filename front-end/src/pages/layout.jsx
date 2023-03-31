import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import { useEffect, useState } from "react";
import Mapshow from '../component/mapshow'
import {useNavigate} from "react-router-dom";
import StickyFooter from "../component/footer";
import HeatmapComponent from "../component/heatmap";
import HeatMapshow from "../component/heatmap";
import Leafheatmap from '../component/leafletheatmap'
import Paper from "@mui/material/Paper";


export  function Layout() {
    const[userlat,setLat] = useState('')
    const[userlng,setLng] = useState('')
    const nav = useNavigate();

    const handleSignup=()=>{
        nav('/signup')
    }
    const handleLogin=()=>{
        nav('/login')
    }
    if (navigator.geolocation){
        navigator.geolocation.getCurrentPosition(pos=>{
            setLat(pos.coords.latitude)
            setLng(pos.coords.longitude)

            // 获取到了用户当前位置的坐标

        },error=>{
            switch(error.code){
                case error.PERMISSION_DENIED:
                    alert("please enable location access");
                    break;
                case error.POSITION_UNAVAILABLE:
                    alert("Fail to locate");
                    break;
                case error.TIMEOUT:
                    alert("Timeout");
                    break;
                default:
                    alert("Fail to locate");
            }
        });
    }
    return (

        <Box sx={{ flexGrow: 1 ,backgroundImage:'url(/sunsetbackground.jpg)',

            backgroundColor: (t) =>
                t.palette.mode === 'light' ? t.palette.grey[50] : t.palette.grey[900],
            backgroundSize: 'cover',
            backgroundPosition: 'center'}}>
            <AppBar position="static">
                <Toolbar>
                    <IconButton
                        size="large"
                        edge="start"
                        color="inherit"
                        aria-label="menu"
                        sx={{ mr: 2 }}
                    >
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                        AfterglowX
                    </Typography>

                    <Button color="inherit" onClick={handleLogin}>login</Button>
                    <Button color="inherit" onClick={handleSignup}>signup</Button>

                </Toolbar>
            </AppBar>
            {/*<Mapshow ></Mapshow>*/}
            {/*<HeatMapshow></HeatMapshow>*/}

            <Leafheatmap userLat={userlat} userLng={userlng}></Leafheatmap>

            <StickyFooter></StickyFooter>
        </Box>

    );
}