import './Form_File.css'

import { Box, Card, TextField, Button, ToggleButtonGroup, FormControl, Select, InputLabel, MenuItem } from '@mui/material'
import ToggleButton from '../components/ToggleButton'
import LockIcon from '@mui/icons-material/Lock'
import LockOpenIcon from '@mui/icons-material/LockOpen'
import DoubleArrowIcon from '@mui/icons-material/DoubleArrow';
import { useState } from 'react'

function Form_File() {
    const [plain, setPlain] = useState('')
    const [cipher, setCipher] = useState('')
    const [mode, setMode] = useState('')
    const [isEncrypt, setBool] = useState(true)

    const handleToggle = (event, newDirection) => {
        if (newDirection !== null) {
            setBool(newDirection === 'true' ? true : false)
        }
    }

    const handleChange = (event) => {
        setMode(event.target.value)
    }


    const handleSubmit = (event) => {
        // TODO: Implement real event handling
        event.preventDefault()


        // TODO: set loading state
        const key = event.target.key.value

        if (isEncrypt) {
            const plaintext = event.target.plaintext.value
            console.log(key, mode, plaintext)
            setCipher('The quick brown fox jumps over the lazy dog.')
        } else {
            const ciphertext = event.target.ciphertext.value
            console.log(key, mode, ciphertext)
            setPlain('The quick brown fox jumps over the lazy dog.')
        }
    }

    return (
        <Box
            sx={{
                display: 'flex',
                // paddingLeft: '2.5rem',
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

export default Form_File