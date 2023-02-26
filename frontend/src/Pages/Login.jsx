import * as React from 'react';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import { ArrowForward } from '@mui/icons-material';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { capitalize } from '@mui/material';
import { MuiTelInput } from 'mui-tel-input'

const theme = createTheme();

export default function Login() {
  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    console.log({
      email: data.get('email'),
      password: data.get('password'),
    });
  };

  const [value, setValue] = React.useState('')

  const handleChange = (newValue) => {
    setValue(newValue)
  }

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="l">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Typography component="h1" variant="h4">
            Регистрация  ученика
          </Typography>
          <Typography fontSize= '12' color='gray'>
             Репетиторы не будут видеть ваши контакты, они нужны только для регистрации
          </Typography>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
            <Grid container spacing={2}>
              
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  variant="standard"
                  name="name"
                  label="Имя"
                  type="name"
                  id="name"
                  autoComplete="Имя"
                />
              </Grid>
              <Grid item paddingTop={10} xs={12} >
              <MuiTelInput
               fullWidth
               langOfCountryName="rus"
               value={value}
               defaultCountry = "RU"
               forceCallingCode
               variant="standard"
               onChange={handleChange} />
               
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  variant="standard"
                  id="email"
                  label="Эл. почта"
                  name="email"
                  autoComplete="Эл. почта"
                />
              </Grid>
              <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2, borderRadius: 15, textTransform: 'capitalize' , justifyContent: 'center'}}
              endIcon={<ArrowForward/>}
            >
              Зарегистрироваться
            </Button>
              <Grid item xs={20} color='grey' >
                <FormControlLabel
                  fontSize='5'
                  control={<Checkbox value="allowExtraEmails" color="primary" />}
                  label='Нажимая "Зарегистрироваться" вы соглашаетесь на обработку персональных данных и условия испльзования сервиса'
                />
              </Grid>
            </Grid>
            <Grid container justifyContent="flex-end">
              <Grid item paddingTop={10}  xs={20}>
                Уже есть аккаунт?
                <Link paddingLeft={2} href="/Tasks" variant="body2">
                  Войти в личный кабинет
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}