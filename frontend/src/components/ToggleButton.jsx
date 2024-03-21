import MuiToggleButton from "@mui/material/ToggleButton";
import { styled } from "@mui/material/styles";

const ToggleButton = styled(MuiToggleButton)(({ theme }) => ({
    "&.MuiToggleButton-root": {
        color: theme.palette.primary.contrastText,
        backgroundColor: theme.palette.primary.main,
        "&.Mui-selected": {
            color: theme.palette.secondary.contrastText,
            backgroundColor: theme.palette.secondary.main,
        },
        "&.Mui-selected:hover": {
            color: theme.palette.secondary.contrastText,
            backgroundColor: theme.palette.secondary.main,
        },
    },
}));

export default ToggleButton;