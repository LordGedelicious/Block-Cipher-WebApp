import React from 'react'
import ReactDOM from 'react-dom/client'
import Form from './partials/Form.jsx'
import './index.css'
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { CssBaseline } from '@mui/material';

import Footer from './partials/Footer.jsx'
import Header from './partials/Header.jsx'

const theme = createTheme({
    palette: {
        // Set mode programatically
        mode: 'dark',
        primary: {
            main: '#93B7BE',
        },
        secondary: {
            main: '#022B3A',
        },
        error: {
            main: '#6F584B',
        },
        warning: {
            main: '#FBB7C0',
        },
        info: {
            main: '#5A352A',
        },
        success: {
            main: '#386641',
        },
    },
});

ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Header />
            <Form />
            <Footer />
        </ThemeProvider>
    </React.StrictMode>,
)
