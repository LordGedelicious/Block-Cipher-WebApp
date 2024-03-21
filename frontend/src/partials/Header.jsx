import './Header.css'
import Paper from '@mui/material/Paper';

function Header() {
    return (
        <header>
            <Paper
                square
                elevation={10}
                sx={{
                    display: "flex",
                    flexDirection: "row",
                    height: "100%",
                    width: "fit-content",
                    alignItems: "center",
                    borderRadius: "0 0 1rem 0",
                    px: "1rem",
                    py: "1rem",
                }}
            >
                <img src="/logo.png" alt="GoBlockC Gopher" className='gopher'/>
                <h1>GoBlock&Xi;</h1>
            </Paper>
        </header>
    )
}

export default Header