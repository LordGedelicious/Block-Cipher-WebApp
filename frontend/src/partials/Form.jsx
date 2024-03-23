import './Form.css'
import Form_Text from './Form_Text';
import Form_File from './Form_File';

import { Container, Box, ToggleButtonGroup } from '@mui/material';
import ToggleButton from '../components/ToggleButton'
import AbcIcon from '@mui/icons-material/Abc'
import DescriptionIcon from '@mui/icons-material/Description'
import { useState } from 'react';

function Form() {
    // On form submit, send a POST request to the server to encrypt/decrypt data
    const fetchBE = async (message, key, encryptOrDecrypt='encrypt', mode='ecb') => {
        const response = await fetch('http://localhost:8080/goblockc', {
            method: 'POST',
            body: JSON.stringify({ 
            message: message,
            key: key,
            encryptOrDecrypt: encryptOrDecrypt,
            mode: mode,
            }),
        })
        const data = await response.json()
        console.log("Results from encrypt:", data)
        return data
    }

    const [mode, setMode] = useState('text')

    const handleModeChange = (event, newMode) => {
        if (newMode !== null) {
            setMode(newMode)
        }
    }

    return (
        <main>
            <Container
                sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                    height: '80vh',
                }}
            >
                <Box
                    sx={{
                        alignItems: 'end',
                        width: "fit-content",
                        height: "100%",
                        display: "flex",
                        justifyContent: "center",
                        flexDirection: "column",
                    }}
                >
                    <ToggleButtonGroup
                        className='mode-toggle'
                        color="warning"
                        value={mode}
                        exclusive
                        onChange={handleModeChange}
                    >
                        <ToggleButton value="text" >
                            <AbcIcon />
                        </ToggleButton>
                        <ToggleButton value="file">
                            <DescriptionIcon />
                        </ToggleButton>
                    </ToggleButtonGroup>
                    {mode === 'text'
                        ? <Form_Text fetchBE={fetchBE} />
                        : <Form_File fetchBE={fetchBE} />
                    }
                </Box>
            </Container>
        </main>
    )
}

export default Form
