import './Footer.css'
import Paper from '@mui/material/Paper';
import Link from '@mui/material/Link';

function Footer() {
    return (
        <footer>
            <Paper
                square
                elevation={10}
                sx={{
                    width: "100%",
                    height: "100%",
                    textAlign: "center",
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                }}
            >
                &copy; 2024 &middot; Gede, Nathan, Felicia &middot;&nbsp;<Link color="inherit" href="https://opensource.org/license/mit" underline="hover">MIT License</Link> 
                {/* TODO: change lisence link to repo */}
            </Paper>
        </footer>
    )
}

export default Footer