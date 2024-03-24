import './Form_File.css'

import { Box, Card, TextField, Button, Switch, Typography, ToggleButtonGroup, Snackbar, FormControl, Select, InputLabel, MenuItem, Alert } from '@mui/material'
import ToggleButton from '../components/ToggleButton'
import VisuallyHiddenInput from '../components/VisuallyHiddenInput'
import LockIcon from '@mui/icons-material/Lock'
import LockOpenIcon from '@mui/icons-material/LockOpen'
import DoubleArrowIcon from '@mui/icons-material/DoubleArrow';
import CloudUploadIcon from '@mui/icons-material/CloudUpload';
import { useState } from 'react'

function Form_File() {
    const [preview, setPreview] = useState('')
    const [result, setResult] = useState('')
    const [filename, setFilename] = useState('')
    const [time, setTime] = useState('')
    const [mode, setMode] = useState('')
    const [isEncrypt, setBool] = useState(true)
    const [open, setOpen] = useState(false)
    const [error, setError] = useState('')

    const handleClose = () => {
        setOpen(false)
    }

    const handleUpload = (event) => {
        const file = event.target.files[0]
        if (event.target.files[0].size > 16777216) {
            setError('File size must be less than 16MiB')
            setOpen(true)
            return
        }
        const reader = new FileReader()
        reader.onload = (event) => {
            var res = [...new Uint8Array(event.target.result)]
                .map(x => x.toString(16).padStart(2, '0'));
            setResult(res.join(''))
            setPreview(res.slice(0, 1024).join(' ') + (res.length > 1024 ? ' .. .. ..' : ''))
            setFilename(file.name)
        }
        reader.readAsArrayBuffer(file)
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
                isFile: true
            }),
        })
        const data = await response.json()

        setResult(data.result)
        console.log(data.result);
        setPreview(data.result.slice(0, 1024).join(' ') + (res.length > 1024 ? ' .. .. ..' : ''))
        setTime(data.timeElapsed)
        var bytes = new Uint8Array(Math.ceil(data.result / 2));
        for (var i = 0; i < bytes.length; i++) {
            bytes[i] = parseInt(data.result.substr(i * 2, i * 2 + 2), 16);
        }
        console.log(bytes);
        var blob = new Blob(bytes, { type: 'application/octet-stream' })
        
        const url = window.URL.createObjectURL(blob)

        const a = document.createElement('a')
        a.href = data
        a.download = filename + (isEncrypt ? '.enc' : '.dec')
        document.body.appendChild(a)
        a.style.display = 'none'
        a.click()
        a.remove()
        setTimeout(() => {
            window.URL.revokeObjectURL(url)
        }, 1000);
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
        var valid = key.toLowerCase().split('').every(c => '0123456789abcdef'.indexOf(c) !== -1);
        if (!valid) {
            setError('Key must be a valid hexadecimal string')
            setOpen(true)
            return
        }

        var message = result
        fetchBE(message, key, true, isEncrypt, mode)
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
                            flexDirection: 'column',
                            width: '100%',
                            // justifyContent: 'space-between',
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
                                id="preview"
                                label="Preview"
                                value={preview}
                                onChange={(event) => { setPreview(event.target.value) }}
                                multiline
                                rows={8}
                                disabled
                                sx={{
                                    margin: '1rem',
                                    flex: 1
                                }}
                                inputProps={{
                                    style: {
                                        fontFamily: 'monospace',
                                    }
                                }}
                            />
                        </Box>
                        <Box
                            sx={{
                                display: 'flex',
                                flexDirection: 'row',
                                margin: '1rem',
                                alignItems: 'center'
                            }}
                        >
                            <Button
                                component="label"
                                variant="contained"
                                tabIndex={-1}
                                startIcon={<CloudUploadIcon />}
                                sx={{
                                    marginRight: '1rem'
                                }}
                            >
                                Upload File
                                <VisuallyHiddenInput
                                    type="file"
                                    onChange={handleUpload}
                                />
                            </Button>
                            <Typography>
                                {filename.length > 36 ? filename.slice(0, 32) + '[...].' + filename.split('.').pop() : filename}
                            </Typography>
                        </Box>
                    </Box>
                    <div style={{ display: 'flex', width: '100%', justifyContent: 'space-between', paddingLeft: '1rem', paddingRight: '1rem' }}>
                        <p>{time !== '' ? "Time elapsed: " + time : ''}</p>
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

export default Form_File