import React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
import Button  from '@mui/material/Button' ;
import cl from './MyModal.module.css'
const style = {
    textAlign: 'center',
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
  };


const MyModal = ({children, visible,titlebox,titlebutton, handleOpen, handleClose,style}) => {

    

    return (
        <div style = {{textAlign: 'center'}}>
      <Button variant="outlined"  onClick={handleOpen}>{titlebutton}</Button>
      <Modal
        open={visible}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            {titlebox}
          </Typography>
          <Typography id="modal-modal-description" sx={{ mt: 2 } }  >
            {children}
          </Typography>
        </Box>
      </Modal>
    </div>
    );
};

export default MyModal;