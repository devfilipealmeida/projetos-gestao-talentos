import { TextField, Button, Container, Typography, Link } from '@mui/material';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useNavigate } from 'react-router-dom';
import api from '../api/axios';

const loginSchema = z.object({
  email: z.string().email(),
  senha: z.string().min(6),
});

type LoginForm = z.infer<typeof loginSchema>;

const Login = () => {
  const navigate = useNavigate();
  const { register, handleSubmit, formState: { errors } } = useForm<LoginForm>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginForm) => {
    try {
      const response = await api.post('/login', data);
      localStorage.setItem('token', response.data.token);
      navigate('/candidatos');
    } catch (error) {
      console.error('Erro ao fazer login', error);
    }
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" gutterBottom>Login</Typography>
      <form onSubmit={handleSubmit(onSubmit)}>
        <TextField {...register('email')} label="Email" fullWidth margin="normal" error={!!errors.email} helperText={errors.email?.message} />
        <TextField {...register('senha')} label="Password" type="password" fullWidth margin="normal" error={!!errors.senha} helperText={errors.senha?.message} />
        <Button type="submit" variant="contained" color="primary" fullWidth>Entrar</Button>
      </form>
      <Typography variant="body2" align="center" style={{ marginTop: '1rem' }}>
        Não tem uma conta? <Link href="/registrar-adm">Registre-se aqui</Link>
      </Typography>
    </Container>
  );
};

export default Login;
