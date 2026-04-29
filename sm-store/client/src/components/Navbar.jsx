import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import {
	Sparkles,
	Search,
	Bell,
	User,
	ShoppingBag,
	ArrowRight,
	Menu,
} from "lucide-react";
import { GradientButton } from '@/components/Buttons';
import { zUser } from '@/store/user';

import { POPPINS, GRADIENT2 } from '@/config/style';

const Navbar = () => {
	const [scrolled, setScrolled] = useState(false);
	const [menuOpen, setMenuOpen] = useState(false);

	const isLoggedIn = zUser(state => state.isLoggedIn);

	useEffect(() => {
		const fn = () => setScrolled(window.scrollY > 20);
		window.addEventListener("scroll", fn);
		return () => window.removeEventListener("scroll", fn);
	}, []);

	return (
		<nav
			className={`fixed top-0 left-0 right-0 z-50 transition-all duration-300 ${scrolled ? "bg-white/90 backdrop-blur-md shadow-sm" : "bg-transparent"
				}`}
			style={{ fontFamily: POPPINS }}
		>
			<div className="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
				{/* Logo */}
				<Link to="/">
					<div className="flex items-center gap-2">
						<div className="w-9 h-9 rounded-xl flex items-center justify-center shadow-md"
							style={{ background: GRADIENT2 }}>
							<Sparkles size={16} className="text-white" />
						</div>
						<span className="text-lg font-bold text-gray-900 tracking-tight">Luminary</span>
					</div>
				</Link>

				{/* Nav links */}
				<div className="hidden md:flex items-center gap-8">
					{["Shop", "Collections", "Deals", "About"].map((link) => (
						<a key={link} href="#"
							className="text-sm font-medium text-gray-600 hover:text-purple-700 transition-colors">
							{link}
						</a>
					))}
				</div>

				{/* Actions */}
				<div className="hidden md:flex items-center gap-3">
					<button className="w-9 h-9 rounded-xl flex items-center justify-center text-gray-500 hover:bg-gray-100 transition-colors">
						<Search size={18} />
					</button>
					{isLoggedIn ? (
						<>
							<button className="w-9 h-9 rounded-xl flex items-center justify-center text-gray-500 hover:bg-gray-100 transition-colors">
								<Bell size={18} />
							</button>
							<Link to="/profile" className="w-9 h-9 rounded-xl flex items-center justify-center text-gray-500 hover:bg-gray-100 transition-colors">
								<User size={18} />
							</Link>
							<button className="relative w-9 h-9 rounded-xl flex items-center justify-center text-gray-500 hover:bg-gray-100 transition-colors">
								<ShoppingBag size={18} />
								<span className="absolute top-1.5 right-1.5 w-2 h-2 rounded-full bg-pink-500" />
							</button>

							<Link to="/shop">
								<GradientButton className="ml-2 cursor-pointer">
									Shop Now <ArrowRight size={14} />
								</GradientButton>
							</Link>
						</>
					) : (
						<Link to="/signin">
							<GradientButton className="ml-2 cursor-pointer">
								Get Started <ArrowRight size={14} />
							</GradientButton>
						</Link>
					)}
				</div>

				{/* Mobile menu toggle */}
				<button onClick={() => setMenuOpen(!menuOpen)}
					className="md:hidden w-9 h-9 flex items-center justify-center text-gray-700">
					{menuOpen ? <X size={20} /> : <Menu size={20} />}
				</button>
			</div>

			{/* Mobile menu */}
			{menuOpen && (
				<div className="md:hidden bg-white border-t border-gray-100 px-6 py-4 flex flex-col gap-4">
					{["Shop", "Collections", "Deals", "About"].map((link) => (
						<a key={link} href="#" className="text-sm font-medium text-gray-700">{link}</a>
					))}
					<GradientButton className="w-full justify-center mt-2">
						Shop Now <ArrowRight size={14} />
					</GradientButton>
				</div>
			)}
		</nav>
	);
}

export default Navbar;