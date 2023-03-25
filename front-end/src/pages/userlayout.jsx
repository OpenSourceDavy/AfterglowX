import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import AccountCircle from '@mui/icons-material/AccountCircle';

import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import Mapshow from "../component/mapshow";
import {useNavigate} from "react-router-dom";
import {useLocation} from "react-router-dom";
import {useEffect, useState} from "react";

import StickyFooter from "../component/footer";
import Leafheatmap from "../component/leafletheatmap";


export default function Userlayout() {
    const[userlat,setLat] = useState('')
    const[userlng,setLng] = useState('')
    const [anchorEl, setAnchorEl] = React.useState(null);
    let nav = useNavigate();
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

    const handleMenu = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };
    const handleNewNotification=()=>{
        nav('/notification')
    }
    const handleNotification=()=>{
        let user_id = location.state.name
        let user_token = location.state.tok
        nav('/curnotification',{state:{name: {user_id}, tok:{user_token}}})
    }
    const handleLogout=()=>{
        nav('/')
    }
    const location = useLocation();

    useEffect(()=>{
        if (location.state == null){
            alert("unauthorized visit")
            nav('/')
        }
        else{const user_id = location.state.name
            const user_token = location.state.tok
        console.log(user_id)
        console.log(user_token);}},[])
    return (
        <Box sx={{ flexGrow: 1,backgroundImage:'url(/sunsetbackground.jpg)',

            backgroundColor: (t) =>
                t.palette.mode === 'light' ? t.palette.grey[50] : t.palette.grey[900],
            backgroundSize: 'cover',
            backgroundPosition: 'center' }}>
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

                        <div>
                            <IconButton
                                size="large"
                                aria-label="account of current user"
                                aria-controls="menu-appbar"
                                aria-haspopup="true"
                                onClick={handleMenu}
                                color="inherit"
                            >
                                <AccountCircle />
                            </IconButton>
                            <Menu
                                id="menu-appbar"
                                anchorEl={anchorEl}
                                anchorOrigin={{
                                    vertical: 'top',
                                    horizontal: 'right',
                                }}
                                keepMounted
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'right',
                                }}
                                open={Boolean(anchorEl)}
                                onClose={handleClose}
                            >
                                <MenuItem onClick={handleNewNotification}>New Notification</MenuItem>
                                <MenuItem onClick={handleNotification}>Current Notification</MenuItem>
                                <MenuItem onClick={handleLogout}>Log out</MenuItem>
                            </Menu>
                        </div>

                </Toolbar>
            </AppBar>


            <Leafheatmap userLat={userlat} userLng={userlng}></Leafheatmap>
            <StickyFooter></StickyFooter>

        </Box>
    );
}