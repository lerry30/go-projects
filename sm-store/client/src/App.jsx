import { useEffect } from 'react';
import { Routes, Route } from 'react-router-dom';
import { zUser } from '@/store/user';

import { apiRequest } from '@/utils/send';
import { BASE_URL } from '@/config/server';

import Home from '@/pages/Home';
import Signin from '@/pages/Signin';
import Signup from '@/pages/Signup';
import Profile from '@/pages/Profile';
import Shop from '@/pages/shop';
import TermsOfService from '@/pages/TermsOfService';
import PrivPolicy from '@/pages/PrivPolicy';
import CukiePolicy from '@/pages/CukiePolicy';
import NotFound from '@/pages/NotFound';

function App() {
	const zSetUserState = zUser(state => state.setUserState);

	useEffect(() => {
		(async () => {
			try {
				const result = await apiRequest(`${BASE_URL}/priv/user`, {auth: true});
				if(result) {
					zSetUserState(true);
				}
			} catch(err){}
		})()
	}, []);

	return (
		<Routes>
			<Route path="/" element={<Home />} />
			<Route path="/signin" element={<Signin />} />
			<Route path="/signup" element={<Signup />} />
			<Route path="/profile" element={<Profile />} />
			<Route path="/shop" element={<Shop />} />
			<Route path="/termsofservice" element={<TermsOfService/>} />
			<Route path="/privpolicy" element={<PrivPolicy />} />
			<Route path="/cukiepolicy" element={<CukiePolicy />} />
			<Route path="*" element={<NotFound />} />
		</Routes>
	);
}

export default App;