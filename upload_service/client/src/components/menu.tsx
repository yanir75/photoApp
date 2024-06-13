import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import MenuItem from '@mui/material/MenuItem';
import AdbIcon from '@mui/icons-material/Adb';

import {  useNavigate} from "react-router-dom";
import { useState } from 'react';


const pages = ['Upload', 'Logout',"Gallery"];

function ResponsiveAppBar(props:any) {
    if(props.visble !== false){
      props.visble = true
    }
    let navigate = useNavigate();

    function navigation(item : any){
      // console.log(item)
      if (item === "/logout"){
        navigate('/user')
      }
      else {
        navigate(item,{state: {keys: props.countries}})
      }
    }
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [anchorElNav, setAnchorElNav] = useState<null | HTMLElement>(null);

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    navigation(`/`+event.currentTarget.id.toLowerCase());
    navigation(0)

    setAnchorElNav(event.currentTarget)
  };


  const handleCloseNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    navigation(`/`+event.currentTarget.id.toLowerCase());
    navigation(0)

    setAnchorElNav(null)
  };
 

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    if (anchorEl === null){
    setAnchorEl(event.currentTarget);
    }
    else{
      setAnchorEl(null);

    }
  };

  const handleClose = (event: React.MouseEvent<HTMLElement>) => {
    // console.log(event.currentTarget.id)
    if (event.currentTarget.id){
      navigation("/country/".concat(event.currentTarget.id))
    }
    setAnchorEl(null);
  };

  return (
    <div style={{display: props.visble == true ? "inline" : "none"}}>
    <AppBar position="fixed">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
        <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
            onClick={handleMenu}

          >
            <MenuIcon/>
            <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'left',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'left',
                }}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                {props.countries ? props.countries.map((country: string)=> <MenuItem id={country} onClick={handleClose}>{country}</MenuItem>) : null}
                {/* <MenuItem  onClick={handleClose}>Profile</MenuItem> */}
                {/* <MenuItem onClick={handleClose}>My account</MenuItem> */}
              </Menu>

          </IconButton>
          <AdbIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
          <Typography
            variant="h6"
            noWrap
            component="a"
            // href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: 'none', md: 'flex' },
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            LOGO
          </Typography>

          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: 'bottom',
                horizontal: 'left',
              }}
              keepMounted
              transformOrigin={{
                vertical: 'top',
                horizontal: 'left',
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: 'block', md: 'none' },
              }}
            >
              {pages.map((page) => (
                <MenuItem key={page} onClick={handleCloseNavMenu}>
                  <Typography textAlign="center">{page}</Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
          <AdbIcon sx={{ display: { xs: 'flex', md: 'none' }, mr: 1 }} />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: 'flex', md: 'none' },
              flexGrow: 1,
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            LOGO
          </Typography>
          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
            {pages.map((page) => (
              <Button
                key={page}
                onClick={handleCloseNavMenu}
                id={page}
                sx={{ my: 2, color: 'white', display: 'block' }}
              >
                {page}
              </Button>
            ))}
          </Box>

        </Toolbar>
      </Container>
    </AppBar>
    </div>
  );
}
export default ResponsiveAppBar;
