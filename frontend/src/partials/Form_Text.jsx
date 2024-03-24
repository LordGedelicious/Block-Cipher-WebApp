import './Form_Text.css'

import { Box, Card, TextField, Button, Switch, Typography, ToggleButtonGroup, Snackbar, FormControl, Select, InputLabel, MenuItem, Alert } from '@mui/material'
import LockIcon from '@mui/icons-material/Lock'
import ToggleButton from '../components/ToggleButton'
import LockOpenIcon from '@mui/icons-material/LockOpen'
import DoubleArrowIcon from '@mui/icons-material/DoubleArrow';
import { useState } from 'react'

function Form_Text() {
    const [plain, setPlain] = useState('')
    const [cipher, setCipher] = useState('')
    const [time, setTime] = useState('')
    const [mode, setMode] = useState('')
    const [isEncrypt, setBool] = useState(true)
    const [isHex, setHex] = useState(false)
    const [open, setOpen] = useState(false)
    const [error, setError] = useState('')

    const handleClose = () => {
        setOpen(false)
    }

    // Send a POST request to the server to encrypt/decrypt data
    const fetchBE = async (message, key, isHex=false, isEncrypt=true, mode='ecb') => {
        const response = await fetch('http://localhost:8080/goblockc', {
            method: 'POST',
            body: JSON.stringify({ 
                message: message,
                isHex: isHex,
                key: key,
                isEncrypt: isEncrypt,
                mode: mode,
            }),
        })
        const data = await response.json()

        if (isEncrypt) {
            setCipher(data.result)
        } else {
            setPlain(data.result)
        }
        setTime(data.timeElapsed)
    }

    const handleToggle = (event, newDirection) => {
        if (newDirection !== null) {
            setBool(newDirection === 'true' ? true : false)
        }
    }

    const handleChange = (event) => {
        setMode(event.target.value)
    }

    const handleHex = (event) => {
        setHex(event.target.checked)
    }

    const handleSubmit = (event) => {
        event.preventDefault()

        // TODO: set loading state

        // TODO: Should key always be hex or string?
        const key = event.target.key.value
        valid = key.toLowerCase().split('').every(c => '0123456789abcdef'.indexOf(c) !== -1);
        if (!valid) {
            setError('Key must be a valid hexadecimal string')
            setOpen(true)
            return
        }

        var message
        var valid
        if (isEncrypt) {
            message = event.target.plaintext.value
            if (isHex) {
                valid = message.toLowerCase().split('').every(c => '0123456789abcdef'.indexOf(c) !== -1);
                if (!valid) {
                    setError('Plaintext must be a valid hexadecimal string')
                    setOpen(true)
                    return
                }
            }
            fetchBE(message, key, isHex, isEncrypt, mode)
        } else {
            message = event.target.ciphertext.value
            valid = message.toLowerCase().split('').every(c => '0123456789abcdef'.indexOf(c) !== -1);
            if (!valid) {
                setError('Ciphertext must be a valid hexadecimal string')
                setOpen(true)
                return
            }
            fetchBE(message, key, true, isEncrypt, mode)
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
                            position: 'relative',
                            display: 'flex',
                            width: '100%',
                            justifyContent: 'space-between',
                            marginBottom: '2rem'
                        }}
                    >
                        <Box
                            sx={{
                                position: 'relative',
                                display: 'flex',
                                flex: 1
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
                            <Box
                                sx={{
                                    position: 'absolute',
                                    bottom: "-1.25rem",
                                    left: "1rem",
                                    display: 'flex',
                                    flexDirection: 'row',
                                    alignItems: 'center',
                                }}
                            >
                                <Typography
                                    sx={{
                                        color: !isEncrypt? 'gray': 'white',
                                    }}
                                >
                                    String
                                </Typography>
                                <Switch
                                    value={isHex}
                                    onChange={handleHex}
                                    disabled={!isEncrypt}
                                />
                                <Typography
                                    sx={{
                                        color: !isEncrypt? 'gray': 'white',
                                    }}
                                >
                                    Hexadecimal
                                </Typography>
                            </Box>
                        </Box>
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
                    <div style={{ display: 'flex', width: '100%', justifyContent: 'space-between', paddingLeft: '1rem',paddingRight: '1rem' }}>
                        <p>Time elapsed: {time}</p>
                        <Button
                            variant='contained'
                            type='submit'
                        >
                            <DoubleArrowIcon />
                            &nbsp;Go!
                        </Button>
                    </div>
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
            <Snackbar anchorOrigin={{ vertical: 'top', horizontal: 'right' }} open={open} autoHideDuration={6000} onClose={handleClose}>
                <Alert
                onClose={handleClose}
                severity="warning"
                variant="filled"
                    sx={{
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        width: '100%' 
                    }}
                >
                    {error}
                </Alert>
            </Snackbar>
        </Box>
    )
}

export default Form_Text