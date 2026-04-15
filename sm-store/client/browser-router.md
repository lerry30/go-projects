Here's how to use `BrowserRouter` in React:

## Installation

```bash
npm install react-router-dom
```

## Basic Setup

Wrap your app with `BrowserRouter` in `index.js` or `App.js`:

```jsx
import { BrowserRouter } from 'react-router-dom';
import App from './App';

ReactDOM.createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <App />
  </BrowserRouter>
);
```

## Defining Routes

```jsx
import { Routes, Route } from 'react-router-dom';

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/about" element={<About />} />
      <Route path="/user/:id" element={<UserProfile />} />
      <Route path="*" element={<NotFound />} />  {/* 404 fallback */}
    </Routes>
  );
}
```

## Navigation

**Using `<Link>` (declarative):**
```jsx
import { Link } from 'react-router-dom';

function Navbar() {
  return (
    <nav>
      <Link to="/">Home</Link>
      <Link to="/about">About</Link>
    </nav>
  );
}
```

**Using `useNavigate` (programmatic):**
```jsx
import { useNavigate } from 'react-router-dom';

function Login() {
  const navigate = useNavigate();

  const handleLogin = () => {
    // after login logic...
    navigate('/dashboard');         // go to route
    navigate(-1);                   // go back
    navigate('/dashboard', { replace: true }); // replace history
  };
}
```

## Reading URL Params & Query Strings

```jsx
import { useParams, useSearchParams } from 'react-router-dom';

// URL: /user/42
function UserProfile() {
  const { id } = useParams(); // id = "42"
}

// URL: /search?query=react
function Search() {
  const [searchParams] = useSearchParams();
  const query = searchParams.get('query'); // "react"
}
```

## Nested Routes

```jsx
function App() {
  return (
    <Routes>
      <Route path="/dashboard" element={<Dashboard />}>
        <Route index element={<DashboardHome />} />       {/* /dashboard */}
        <Route path="settings" element={<Settings />} />  {/* /dashboard/settings */}
      </Route>
    </Routes>
  );
}

// In Dashboard.jsx, use <Outlet /> to render child routes
import { Outlet } from 'react-router-dom';

function Dashboard() {
  return (
    <div>
      <Sidebar />
      <Outlet /> {/* Child route renders here */}
    </div>
  );
}
```

## Key Hooks Summary

| Hook | Purpose |
|---|---|
| `useNavigate` | Programmatic navigation |
| `useParams` | Read URL parameters (`:id`) |
| `useSearchParams` | Read query strings (`?key=value`) |
| `useLocation` | Access current location object |
| `useMatch` | Check if path matches current URL |

> **Note:** This uses **React Router v6** syntax. If you're on v5, `<Switch>` replaces `<Routes>` and `useHistory` replaces `useNavigate`.