import './Form_Text.css'

import { Box, Card, TextField, Button, ToggleButtonGroup, FormControl, Select, InputLabel, MenuItem } from '@mui/material'
import LockIcon from '@mui/icons-material/Lock'
import ToggleButton from '../components/ToggleButton'
import LockOpenIcon from '@mui/icons-material/LockOpen'
import DoubleArrowIcon from '@mui/icons-material/DoubleArrow';
import { useState } from 'react'

function Form_Text() {
    const [plain, setPlain] = useState('')
    const [cipher, setCipher] = useState('')
    const [mode, setMode] = useState('')
    const [isEncrypt, setBool] = useState(true)

    // Send a POST request to the server to encrypt/decrypt data
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
        if (encryptOrDecrypt === 'encrypt') {
            setCipher(data)
        } else {
            setPlain(data)
        }
    }

    const handleToggle = (event, newDirection) => {
        if (newDirection !== null) {
            setBool(newDirection === 'true' ? true : false)
        }
    }

    const handleChange = (event) => {
        setMode(event.target.value)
    }

    const handleSubmit = (event) => {
        event.preventDefault()

        // TODO: set loading state
        const key = event.target.key.value

        if (isEncrypt) {
            const plaintext = event.target.plaintext.value
            fetchBE(plaintext, key, 'encrypt', mode)
        } else {
            const ciphertext = event.target.ciphertext.value
            fetchBE(ciphertext, key, 'decrypt', mode)
        }
    }

    return (
        <Box
            sx={{
                display: 'flex',
            }}
        >
            <Card
                className='mode-container'
                sx={{
                    position: 'relative',
                    padding: '1rem'
                }}
            >
                <form
                    onSubmit={handleSubmit}
                    className='text-form'
                >
                    <Box
                        className='input-wrappers'
                        sx={{
                            width: '100%',
                            display: 'flex',
                        }}
                    >
                        <TextField
                            id="key"
                            label="Key"
                            sx={{
                                margin: '1rem',
                                flex: 1,
                            }}
                            required
                        />
                        <FormControl
                            sx={{
                                margin: '1rem',
                                flex: 0.5
                            }}
                        >
                            <InputLabel id="mode-label" required>Mode</InputLabel>
                            <Select
                                labelId="mode-label"
                                value={mode}
                                label="Mode"
                                onChange={handleChange}
                                required
                            >
                                <MenuItem value={'ecb'}>ECB</MenuItem>
                                <MenuItem value={'cbc'}>CBC</MenuItem>
                                <MenuItem value={'ofb'}>OFB</MenuItem>
                                <MenuItem value={'cfb'}>CFB</MenuItem>
                                <MenuItem value={'ctr'}>Counter</MenuItem>
                            </Select>
                        </FormControl>
                    </Box>

                    <Box
                        className='input-wrappers'
                        sx={{
                            display: 'flex',
                            width: '100%',
                            justifyContent: 'space-between',
                            marginBottom: '4rem'
                        }}
                    >
                        <TextField
                            id="plaintext"
                            label="Plaintext"
                            value={plain}
                            onChange={(event) => { setPlain(event.target.value) }}
                            multiline
                            rows={8}
                            disabled={!isEncrypt}
                            required={isEncrypt}
                            sx={{
                                margin: '1rem',
                                flex: 1
                            }}
                        />
                        <TextField
                            id="ciphertext"
                            label="Ciphertext"
                            value={cipher}
                            onChange={(event) => { setCipher(event.target.value) }}
                            multiline
                            rows={8}
                            disabled={isEncrypt}
                            required={!isEncrypt}
                            sx={{
                                margin: '1rem',
                                flex: 1
                            }}
                        />
                    </Box>
                    <Button
                        variant='contained'
                        type='submit'
                        sx={{
                            position: 'absolute',
                            right: 0,
                            bottom: 0,
                            margin: '1rem',
                        }}
                    >
                        <DoubleArrowIcon />
                        &nbsp;Go!
                    </Button>
                </form>
            </Card>
            <ToggleButtonGroup
                className='encrypt-toggle'
                color='primary'
                value={isEncrypt.toString()}
                exclusive
                onChange={handleToggle}
                orientation='vertical'
            >
                <ToggleButton
                    value="true"
                    sx={{
                        zIndex: "inherit",
                    }}
                >
                    <LockIcon />
                </ToggleButton>
                <ToggleButton value="false"
                    sx={{
                        zIndex: "inherit",
                    }}
                >
                    <LockOpenIcon />
                </ToggleButton>
            </ToggleButtonGroup>
        </Box>
    )
}

export default Form_Text