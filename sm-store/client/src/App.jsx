import { Routes, Route } from 'react-router-dom';

import Home from '@/pages/Home';
import Signin from '@/pages/Signin';
import Signup from '@/pages/Signup';
import Profile from '@/pages/Profile';
import TermsOfService from '@/pages/TermsOfService';
import PrivPolicy from '@/pages/PrivPolicy';
import NotFound from '@/pages/NotFound';

function App() {
	return (
		<Routes>
			<Route path="/" element={<Home />} />
			<Route path="/signin" element={<Signin />} />
			<Route path="/signup" element={<Signup />} />
			<Route path="/profile" element={<Profile />} />
			<Route path="/termsofservice" element={<TermsOfService/>} />
			<Route path="/privpolicy" element={<PrivPolicy />} />
			<Route path="*" element={<NotFound />} />
		</Routes>
	);
}

export default App;