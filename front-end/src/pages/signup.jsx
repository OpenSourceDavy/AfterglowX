import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import backURL from '../sunsetBackground.jpg'
import {useState} from "react";
import {useNavigate} from "react-router-dom";

function Copyright(props) {
    return (
        <Typography variant="body2" color="text.secondary" align="center" {...props}>
            {'Copyright © '}
            <Link color="inherit" href="https://afterglow.canadacentral.cloudapp.azure.com">
                afterglow.canadacentral.cloudapp.azure.com
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
}

const theme = createTheme();

export default function SignupLayout() {
    const[name,setName]=useState('')
    const [confirmPassword, setConfirmPassword] = useState("");
    const [passwordErrorText, setPasswordErrorText] = React.useState("");
    const[password,setPassword]=useState('')
    const[email,setEmail]=useState('')
    const[country,setCountry]=useState('')
    const[cell,setCell]=useState('')
    const [nameerror, setnameError] = useState("");
    const [nameCheckError, setnameCheckError] = useState("");
    const [passerror, setpassError] = useState("");
    const [mailerror, setmailError] = useState("");
    const [countryerror, setcountryError] = useState("");
    const [cellerror, setcellError] = useState("");
    const[user, setUser]=useState('')
    const nav = useNavigate();
    const checkPasswords = (event) => {
        if (password !== confirmPassword) {

            setPasswordErrorText("Passwords do not match");
            event.preventDefault();
        } else {
            setPasswordErrorText("");
        }
    }

    // const checkUsername = (event) => {
    //     fetch("/users/isUsername",{
    //         method:"POST",
    //         headers:{"Content-Type":"application/json"},
    //         body:JSON.stringify(name)
    //
    //     }).then(res=>res.json())
    //         .then((res)=>{
    //         if(!res.success){
    //             setnameCheckError("not avaliable")
    //             event.preventDefault();
    //         }
    //         else{
    //             setnameCheckError("avaliable")
    //         }
    //     })
    // }
    const checkUsername = (event) => {
        setnameCheckError("not avaliable")
    }

    const handleSubmit = (e) => {
        checkPasswords(e)
        e.preventDefault();
        if (!email) {
            setmailError("Please enter email");
        } else {
            setmailError("");
        }
        if (!password) {
            setpassError("Please enter password");
        } else {
            setpassError("");
        }
        if (!name) {
            setnameError("Please enter name");
        } else {
            setnameError("");
        }
        if (!cell) {
            setcellError("Please enter cell");
        } else {
            setcellError("");
        }
        if (!country) {
            setcountryError("Please enter country");
        } else {
            setcountryError("");
        }
        const user={name,country,cell,email,password}

        console.log(user)
        fetch("http://ec2-3-83-223-34.compute-1.amazonaws.com:8800/api/v1/register",{
            method:"POST",
            headers:{"Content-Type":"application/json"},
            body:JSON.stringify(user)

        }).then(res=>res.json())
            .then((res)=>{
            if (res){alert("register successfully");
                alert("signup successfully, now redirect to home page")
                nav('/')}

        })

    }


    return (
        <ThemeProvider theme={theme}>
            <Grid container component="main" sx={{ height: '100vh' }}>
                <CssBaseline />
                <Grid
                    item
                    xs={false}
                    sm={4}
                    md={7}
                    sx={{
                        backgroundImage: 'url(/omega-sunset-lymeregis.jpg)',
                        backgroundRepeat: 'no-repeat',
                        backgroundColor: (t) =>
                            t.palette.mode === 'light' ? t.palette.grey[50] : t.palette.grey[900],
                        backgroundSize: 'cover',
                        backgroundPosition: 'center',
                    }}
                />
                <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
                    <Box
                        sx={{
                            my: 8,
                            mx: 4,
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                        }}
                    >
                        <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                            <LockOutlinedIcon />
                        </Avatar>
                        <Typography component="h1" variant="h5">
                            Sign Up
                        </Typography>
                        <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 1 }}>
                            <TextField
                                margin="normal"

                                fullWidth
                                id="Name"
                                label="Name"
                                name="Name"
                                autoComplete="name"
                                required={true}
                                value={name}
                                autoFocus={true}

                                onChange={(e)=>{setName(e.target.value)}}
                                error={nameerror}
                                helperText={nameerror}
                            />
                            <div style={{ color: "red" }}> <Button
                                halfWidth
                                variant="contained"
                                sx={{ mt: 1, mb: 0 }}
                                onClick={(e)=>checkUsername(e)}
                            >
                                Check name availability:  {nameCheckError}
                            </Button></div>
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="password"
                                label="Password"
                                type="password"
                                id="password"
                                value={password}
                                onChange={(e)=>setPassword(e.target.value)}
                                error={passerror}
                                helperText={passerror}
                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="Confirm-Password"
                                label="Confirm-Password"
                                type="Password"
                                id="Confirm-Password"
                                value={confirmPassword}
                                onChange={(e)=>setConfirmPassword(e.target.value)}
                                error={passwordErrorText}
                                helperText={passwordErrorText}

                            />

                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="Email-Address"
                                label="Email-Address"
                                type="Email-Address"
                                id="Email-Address"
                                autoComplete="email"
                                value={email}
                                onChange={(e)=>setEmail(e.target.value)}
                                error={mailerror}
                                helperText={mailerror}

                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name=""
                                label="Country"
                                type="Country"
                                id="Country"
                                autoComplete="country"
                                value={country}
                                onChange={(e)=>setCountry(e.target.value)}
                                error={countryerror}
                                helperText={countryerror}
                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="Cell"
                                label="Cell"
                                type="Cell"
                                id="Cell"
                                autoComplete="tel"
                                value={cell}
                                onChange={(e)=>setCell(e.target.value)}
                                error={cellerror}
                                helperText={cellerror}
                            />
                            <FormControlLabel
                                control={<Checkbox value="remember" color="primary" />}
                                label="Remember me"
                            />
                            <Button
                                type="submit"
                                fullWidth
                                variant="contained"
                                sx={{ mt: 3, mb: 2 }}
                            >
                                Sign Up
                            </Button>
                            <Grid container>
                                <Grid item xs>
                                    <Link href="/" variant="body2">
                                        Forgot password?
                                    </Link>
                                </Grid>
                                <Grid item >
                                    <Link href="/login" variant="body2">
                                        Already Have an Account?
                                    </Link>
                                </Grid>
                            </Grid>
                            <Copyright sx={{ mt: 5 }} />
                        </Box>
                    </Box>
                </Grid>
            </Grid>
        </ThemeProvider>
    );
}