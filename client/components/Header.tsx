"use client";

import Link from "next/link";
import {
  AppBar,
  Box,
  Toolbar,
  Typography,
  Button,
  Divider,
  ListItem,
  ListItemButton,
  List,
  IconButton,
  Drawer,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import { useState } from "react";
import { useAuth } from "@/providers/auth";

const drawerWidth = 240;

export default function Header() {
  const { currentUser } = useAuth();
  const [mobileOpen, setMobileOpen] = useState(false);

  const links = [
    !currentUser && {
      label: "Sign Up",
      href: "/auth/signup",
    },
    !currentUser && {
      label: "Sign In",
      href: "/auth/signin",
    },
    currentUser && {
      label: "Sell Tickets",
      href: "/tickets/new",
    },
    currentUser && {
      label: "My Orders",
      href: "/orders",
    },
    currentUser && {
      label: "Sign Out",
      href: "/auth/signout",
    },
  ].filter(Boolean);

  const handleDrawerToggle = () => {
    setMobileOpen((prevState) => !prevState);
  };

  const drawer = (
    <Box onClick={handleDrawerToggle} sx={{ textAlign: "center" }}>
      <Typography component={Link} href={"/"} variant="h6" sx={{ my: 2 }}>
        Ticket Commerce
      </Typography>
      <Divider />
      <List>
        {links.map((linkConfig) => {
          if (!linkConfig) return null;
          const { label, href } = linkConfig;
          return (
            <ListItem key={href} disablePadding>
              <ListItemButton LinkComponent={Link} href={href} sx={{ textAlign: "center" }}>
                {label}
              </ListItemButton>
            </ListItem>
          );
        })}
      </List>
    </Box>
  );

  return (
    <>
      <AppBar component="nav">
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="start"
            onClick={handleDrawerToggle}
            sx={{ mr: 2, display: { md: "none" } }}
          >
            <MenuIcon />
          </IconButton>
          <Typography
            variant="h6"
            component={Link}
            href="/"
            sx={{ flexGrow: 1, display: { xs: "none", sm: "block" } }}
          >
            Ticket Commerce
          </Typography>
          <Box sx={{ display: { xs: "none", sm: "block" } }}>
            {links.map((linkConfig) => {
              if (!linkConfig) return null;
              const { label, href } = linkConfig;
              return (
                <Button key={href} LinkComponent={Link} href={href} sx={{ color: "#fff" }}>
                  {label}
                </Button>
              );
            })}
          </Box>
        </Toolbar>
      </AppBar>
      <nav>
        <Drawer
          variant="temporary"
          open={mobileOpen}
          onClose={handleDrawerToggle}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
          sx={{
            display: { xs: "block", sm: "none" },
            "& .MuiDrawer-paper": { boxSizing: "border-box", width: drawerWidth },
          }}
        >
          {drawer}
        </Drawer>
      </nav>
    </>
  );
}
