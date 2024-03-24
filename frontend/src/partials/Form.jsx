import './Form.css'
import Form_Text from './Form_Text';
import Form_File from './Form_File';

import { Container, Box, ToggleButtonGroup } from '@mui/material';
import ToggleButton from '../components/ToggleButton'
import AbcIcon from '@mui/icons-material/Abc'
import DescriptionIcon from '@mui/icons-material/Description'
import { useState } from 'react';

function Form() {
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
                        ? <Form_Text />
                        : <Form_File />
                    }
                </Box>
            </Container>
        </main>
    )
}

export default Form
