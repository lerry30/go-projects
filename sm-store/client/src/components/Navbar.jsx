import { Link } from 'react-router-dom';

function Navbar() {
	return (
		<nav>
			<Link to="/">Home</Link>
			<Link to="/profile">Profile</Link>
			<Link to="/signin">SignIn</Link>
			<Link to="/signup">SignUp</Link>
		</nav>
	);
}

export default Navbar;