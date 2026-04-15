import { useState } from 'react';
import { Eye, EyeClosed } from 'lucide-react';
import { Link, useNavigate } from 'react-router-dom';
import { sendJSON } from '@/utils/send';
import { backendUrl } from '@/config/server';
import { SuccessModal, ErrorModal } from '@/components/Modal';

import IllustrationSVG from '@/components/IllustrationSVG';
import Loading from '@/components/Loading';

const StarIcon = () => (
	<svg width="14" height="14" viewBox="0 0 24 24" fill="white">
		<path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" />
	</svg>
);

const EyeIcon = ({ show }) => {
	return show ? <Eye /> : <EyeClosed />
};

const Signup = () => {
	const [form, setForm] = useState({ first_name: "", last_name: "", email: "", phone: "", password: "", confirm: "" });
	const [showPassword, setShowPassword] = useState(false);
	const [showConfirm, setShowConfirm] = useState(false);
	const [agreed, setAgreed] = useState(false);
	const [loading, setLoading] = useState(false);
	const [modal, setModal] = useState({message: '', type: null});

	const navigate = useNavigate()

	const update = (field) => (e) => setForm({ ...form, [field]: e.target.value });

	const passwordStrength = () => {
		const p = form.password;
		if (!p) return null;
		if (p.length < 6) return { label: "Weak", color: "#ef4444", width: "25%" };
		if (p.length < 10) return { label: "Fair", color: "#f97316", width: "55%" };
		return { label: "Strong", color: "#22c55e", width: "100%" };
	};

	const strength = passwordStrength();

	const handleSubmit = async (e) => {
		try {
			e.preventDefault();
			if(!agreed) throw new Error('Read terms and condition first');
			if(!form.password || form.password !== form.confirm) 
				throw new Error('Re-enter your password to confirm');
			
			const tForm = {
				first_name: String(form.first_name).trim(), 
				last_name: String(form.last_name).trim(), 
				email: String(form.email).trim(),
				phone: String(form.phone).trim(),
				password: form.password 
			};

			if(!tForm.first_name || !tForm.last_name || !tForm.email || !tForm.phone) 
				throw new Error('Fields are empty');
			
			setLoading(true);
			const result = await sendJSON(`${backendUrl}/signup`, tForm);
			if(result) {
				setModal({message: 'Account created successfully.', type: 'success'})
				setTimeout(() => {
					setModal({message: '', type: null});
					navigate('/')
				}, 1000);
			}
		} catch(err) {
			console.log('Signup handling form submition error: ', err);
			setModal({message: 'Something went wrong', type: 'error'});
			setTimeout(() => {
				setModal({message: '', type: null});
			}, 1000);
		} finally {
			setLoading(false);
		}
	};

	const fields = [
		{ id: "first_name", label: "First Name", type: "text", placeholder: "Jane" },
		{ id: "last_name", label: "Last Name", type: "text", placeholder: "Doe" },
		{ id: "email", label: "Email Address", type: "email", placeholder: "you@company.com" },
		{ id: "phone", label: "Phone Number", type: "text", placeholder: "+63-912-345-6790"}
	];

	return (
		<div
			className="min-h-screen flex items-center justify-center bg-gray-100 p-4"
			style={{ fontFamily: "'Poppins', sans-serif" }}
		>
			<div className="flex w-full max-w-5xl rounded-3xl overflow-hidden shadow-2xl bg-white"
				style={{ minHeight: "640px" }}>

				{/* ── LEFT PANEL ── */}
				<div className="flex flex-col justify-center w-full lg:w-1/2 px-10 py-12 bg-white">

					{/* Logo */}
					<div className="flex items-center gap-2 mb-10">
						<div
							className="w-8 h-8 rounded-xl flex items-center justify-center"
							style={{ background: "linear-gradient(135deg, #7C3AED, #C026D3)" }}
						>
							<StarIcon />
						</div>
						<span className="text-base font-semibold text-gray-900 tracking-tight">Luminary</span>
					</div>

					{/* Heading */}
					<h1 className="text-3xl font-bold text-gray-900 leading-tight tracking-tight mb-1">
						Create your account
					</h1>
					<p className="text-sm text-gray-400 mb-8">
						Sign up to discover great deals and exclusive offers.
					</p>

					<form onSubmit={handleSubmit} className="space-y-5">

						{/* Name + Email */}
						{fields.map(({ id, label, type, placeholder }) => (
							<div key={id}>
								<label className="block text-xs font-semibold tracking-widest text-gray-400 uppercase mb-2">
									{label}
								</label>
								<input
									type={type}
									value={form[id]}
									onChange={update(id)}
									placeholder={placeholder}
									required
									className="w-full border-0 border-b border-gray-200 pb-2 pt-1 text-sm text-gray-900 bg-transparent outline-none placeholder-gray-300 focus:border-purple-600 transition-colors duration-200"
									style={{ fontFamily: "'Poppins', sans-serif" }}
								/>
							</div>
						))}

						{/* Password */}
						<div>
							<label className="block text-xs font-semibold tracking-widest text-gray-400 uppercase mb-2">
								Password
							</label>
							<div className="relative">
								<input
									type={showPassword ? "text" : "password"}
									value={form.password}
									onChange={update("password")}
									placeholder="Min. 8 characters"
									required
									className="w-full border-0 border-b border-gray-200 pb-2 pt-1 text-sm text-gray-900 bg-transparent outline-none placeholder-gray-300 focus:border-purple-600 transition-colors duration-200 pr-8"
									style={{ fontFamily: "'Poppins', sans-serif" }}
								/>
								<button
									type="button"
									onClick={() => setShowPassword(!showPassword)}
									className="absolute right-0 top-1 text-gray-300 hover:text-purple-500 transition-colors"
								>
									<EyeIcon show={showPassword} />
								</button>
							</div>
							{/* Strength indicator */}
							{strength && (
								<div className="mt-2 flex items-center gap-2">
									<div className="flex-1 h-1 bg-gray-100 rounded-full overflow-hidden">
										<div
											className="h-full rounded-full transition-all duration-300"
											style={{ width: strength.width, background: strength.color }}
										/>
									</div>
									<span className="text-xs font-medium" style={{ color: strength.color }}>
										{strength.label}
									</span>
								</div>
							)}
						</div>

						{/* Confirm Password */}
						<div>
							<label className="block text-xs font-semibold tracking-widest text-gray-400 uppercase mb-2">
								Confirm Password
							</label>
							<div className="relative">
								<input
									type={showConfirm ? "text" : "password"}
									value={form.confirm}
									onChange={update("confirm")}
									placeholder="Re-enter your password"
									required
									className="w-full border-0 border-b border-gray-200 pb-2 pt-1 text-sm text-gray-900 bg-transparent outline-none placeholder-gray-300 focus:border-purple-600 transition-colors duration-200 pr-8"
									style={{ fontFamily: "'Poppins', sans-serif" }}
								/>
								<button
									type="button"
									onClick={() => setShowConfirm(!showConfirm)}
									className="absolute right-0 top-1 text-gray-300 hover:text-purple-500 transition-colors"
								>
									<EyeIcon show={showConfirm} />
								</button>
							</div>
							{form.confirm && form.password !== form.confirm && (
								<p className="text-xs text-red-400 mt-1.5">Passwords don't match</p>
							)}
							{form.confirm && form.password === form.confirm && form.confirm.length > 0 && (
								<p className="text-xs text-green-500 mt-1.5">✓ Passwords match</p>
							)}
						</div>

						{/* Terms */}
						<label className="flex items-start gap-2.5 cursor-pointer">
							<input
								type="checkbox"
								checked={agreed}
								onChange={(e) => setAgreed(e.target.checked)}
								className="accent-purple-600 w-3.5 h-3.5 mt-0.5 shrink-0"
							/>
							<span className="text-xs text-gray-400 leading-relaxed">
								I agree to the{" "}
								<a href="#" className="text-purple-600 font-semibold hover:underline">Terms of Service</a>
								{" "}and{" "}
								<a href="#" className="text-purple-600 font-semibold hover:underline">Privacy Policy</a>
							</span>
						</label>

						{/* CTA */}
						<button
							type="submit"
							disabled={loading || !agreed}
							className="w-full py-3.5 rounded-2xl text-white text-sm font-semibold tracking-wide transition-all duration-200 hover:opacity-90 hover:-translate-y-0.5 active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed disabled:translate-y-0 cursor-pointer"
							style={{
								background: "linear-gradient(135deg, #7C3AED 0%, #C026D3 60%, #f97316 120%)",
								boxShadow: "0 8px 24px rgba(124,58,237,0.28)",
								fontFamily: "'Poppins', sans-serif",
							}}
						>
							{loading ? "Creating your account..." : "Get started for free →"}
						</button>
					</form>

					{/* Sign in link */}
					<p className="text-center text-xs text-gray-400 mt-6">
						Already have an account?{" "}
						<Link to="/signin" className="text-purple-600 font-semibold hover:text-purple-700 transition-colors">
							Sign in →
						</Link>
					</p>
				</div>

				{/* ── RIGHT PANEL ── */}
				<div
					className="hidden lg:flex lg:w-1/2 relative overflow-hidden"
					style={{
						background: "linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%)",
					}}
				>
					{/* Blobs */}
					<div className="absolute -top-16 -right-16 w-80 h-80 rounded-full"
						style={{ background: "rgba(255,255,255,0.07)" }} />
					<div className="absolute -bottom-20 -left-20 w-96 h-96 rounded-full"
						style={{ background: "rgba(255,150,50,0.12)" }} />
					<div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-64 h-64 rounded-full"
						style={{ background: "rgba(255,255,255,0.05)" }} />


					{/* Content */}
					<div className="relative z-10 flex flex-col justify-center px-12 py-14">

						{/* Badge */}
						<div
							className="inline-flex items-center gap-2 self-start rounded-full px-4 py-1.5 mb-8 text-xs font-medium text-white/90 border border-white/20"
							style={{ background: "rgba(255,255,255,0.12)", backdropFilter: "blur(4px)" }}
						>
							<span className="w-1.5 h-1.5 rounded-full bg-green-400" />
							Free delivery on orders over ₱999
						</div>

						<h2 className="text-2xl font-bold text-white leading-snug tracking-tight mb-3">
							Shop smarter,<br />save bigger.
						</h2>
						<p className="text-sm text-white/70 leading-relaxed mb-8">
							Discover thousands of products, get same-day delivery, and enjoy
							hassle-free returns — all backed by our buyer protection guarantee.
						</p>

						<IllustrationSVG />

						{/* Stat cards */}
						<div className="flex gap-3 mt-7">
							{[
							{ num: "500K+", label: "Products listed" },
							{ num: "1-Day", label: "Avg. delivery" },
							{ num: "30-day", label: "Free returns" },
							].map(({ num, label }) => (
							<div
								key={label}
								className="flex-1 rounded-2xl px-4 py-3 border border-white/20"
								style={{ background: "rgba(255,255,255,0.1)", backdropFilter: "blur(6px)" }}
							>
								<div className="text-xl font-bold text-white tracking-tight">{num}</div>
								<div className="text-xs text-white/60 mt-0.5">{label}</div>
							</div>
							))}
						</div>
					</div>
				</div>
			</div>

			{ loading && <Loading /> }

			{(modal?.message && modal?.type === 'success') && (
				<SuccessModal 
					message={modal.message} 
					callback={() => {
						setModal({message: '', type: null});
						navigate('/');
					}} 
				/>
			)}

			{(modal?.message && modal?.type === 'error') && (
				<ErrorModal 
					header="Oops!"
					message={modal.message} 
					callback={() => {
						setModal({message: '', type: null});
					}} 
				/>
			)}
		</div>
	);
}

export default Signup;