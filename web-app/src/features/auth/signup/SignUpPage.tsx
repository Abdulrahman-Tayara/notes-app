import React, {useCallback} from "react";
import useViewModel from "./viewmodel";
import {Avatar, Box, Button, Container, CssBaseline, Grid, Link, TextField, Typography} from "@mui/material";
import {LockOutlined} from "@mui/icons-material";
import {Routes} from "common/routes/constants";

const SignUpPage = () => {
    const {name, email, password, signUp} = useViewModel();

    const handleSubmit = useCallback(() => {
        signUp()
    }, [signUp]);

    return (
        <Container component="main" style={{alignItems: 'center', display: 'flex', flexDirection: 'column'}}>
            <CssBaseline/>
            <Box sx={{
                padding: 8,
                marginTop: 8,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                boxShadow: 10,
                width: "50%",
                borderRadius: '0.5rem'
            }}>
                <Avatar sx={{bgcolor: 'secondary.main'}}>
                    <LockOutlined/>
                </Avatar>

                <Typography component="h1" variant="h5" sx={{marginTop: '5%'}}>
                    Sign up
                </Typography>

                <TextField
                    margin="normal"
                    required
                    fullWidth
                    id="email"
                    label="Email Address"
                    name="email"
                    type="email"
                    autoComplete="email"
                    autoFocus/>

                <TextField
                    margin="normal"
                    required
                    fullWidth
                    id="password"
                    label="Password"
                    name="password"
                    type="password"
                    autoComplete="current-password"/>

                <Button type='submit'
                        fullWidth
                        sx={{marginTop: 8}}
                        variant='contained'>
                    Sign Up
                </Button>


                <Link href={Routes.Login} sx={{marginTop: '5%'}}>
                    Already have an account?
                </Link>

            </Box>
        </Container>
    );
};


export default SignUpPage;