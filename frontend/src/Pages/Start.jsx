import React from "react";
import Login from "./Login";
import MyModal from "../components/UI/MyModal/MyModal";

const style = {
    textAlign: 'center',
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 800,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    borderRadius: 15,
    boxShadow: 24,
    p: 4,
  };
const Start=()=> {

    const [open, setOpen] = React.useState(false);
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);



    return (
        <div>
            <MyModal visible={open} style={style} titlebutton={'Войти'} handleOpen={handleOpen} handleClose={handleClose}>
        <Login />
      </MyModal>
        </div>
    )
} 
export default Start;