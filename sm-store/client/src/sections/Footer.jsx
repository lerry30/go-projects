import {
	Sparkles,
} from "lucide-react";

import { Link } from 'react-router-dom';
import { POPPINS } from '@/config/style';

// ─── Config ─────────────────────────────────────────────────────────────────
const GRADIENT = "linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%)";

const Footer = () => {
	const copyrightLinks = {
		"Privacy Policy": "/privpolicy", 
		"Terms of Service": "/termsofservice", 
		"Cookie Policy": "/cukiepolicy",
	};

	return (
		<footer className="bg-gray-900 text-white pt-16 pb-8" style={{ fontFamily: POPPINS }}>
			<div className="max-w-7xl mx-auto px-6 grid sm:grid-cols-2 lg:grid-cols-4 gap-10 mb-12">
				{/* Brand */}
				<div>
					<div className="flex items-center gap-2 mb-4">
						<div className="w-8 h-8 rounded-xl flex items-center justify-center" style={{ background: GRADIENT }}>
							<Sparkles size={14} className="text-white" />
						</div>
						<span className="text-base font-bold tracking-tight">Luminary</span>
					</div>
					<p className="text-sm text-gray-400 leading-relaxed">Shop smarter. Live brighter. Your one-stop destination for everything you love.</p>
				</div>

				{/* Links */}
				{[
					{ title: "Shop", links: ["New Arrivals", "Best Sellers", "Sale", "Brands"] },
					{ title: "Support", links: ["Help Center", "Track Order", "Returns", "Contact Us"] },
					{ title: "Company", links: ["About", "Careers", "Blog", "Press"] },
				].map((col) => (
					<div key={col.title}>
						<p className="text-xs font-semibold uppercase tracking-widest text-gray-500 mb-4">{col.title}</p>
						<ul className="space-y-2">
							{col.links.map((l) => (
								<li key={l}><a href="#" className="text-sm text-gray-400 hover:text-white transition-colors">{l}</a></li>
							))}
						</ul>
					</div>
				))}
			</div>

			<div className="max-w-7xl mx-auto px-6 pt-8 border-t border-gray-800 flex flex-col sm:flex-row items-center justify-between gap-4">
				<p className="text-xs text-gray-500">© {new Date().getFullYear()} Luminary. All rights reserved.</p>
				<div className="flex gap-6">
					{Object.entries(copyrightLinks).map((l) => {
						//console.log(l[1]);
						return (
						<Link 
							key={l[0]} 
							to={l[1]} 
							className="text-xs text-gray-500 hover:text-white transition-colors"
						>
							{l[0]}
						</Link>
					)})}
				</div>
			</div>
		</footer>
	);
}

export default Footer;