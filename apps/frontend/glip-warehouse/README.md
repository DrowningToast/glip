# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Setup

create `.env` file at `apps/frontend/glip-warehouse` look like this.

```sh
VITE_API_URL=http://localhost:3000
```

then, run this command.

```sh
cd apps/frontend/glip-warehouse
bun install
bun dev
```
