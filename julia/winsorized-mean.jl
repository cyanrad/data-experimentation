### A Pluto.jl notebook ###
# v0.19.29

using Markdown
using InteractiveUtils

# ╔═╡ 12520694-4a43-4074-a959-0c6b2fcb52e0
numbers = [70, 72, 74, 68, 69, 87, 82, 73, 67, 70]

# ╔═╡ 30fa4d88-e32b-46a9-ad01-b87095ba223e
sorted_numbers = sort(numbers)

# ╔═╡ 020f2730-586f-4375-8153-92e1a1faeb41
first_quartile_index = Int(ceil(length(sorted_numbers) / 4))

# ╔═╡ f8bc41b1-25d6-4d68-a907-c24c05a993d2
third_quartile_index = Int(ceil(length(sorted_numbers)*3 / 4))

# ╔═╡ de123813-4c9c-430a-9677-f9cb7af2ac98
first_quartile_value = sorted_numbers[first_quartile_index]

# ╔═╡ f2b31330-244f-40fe-ba97-0d5bd18408a4
third_quartile_value = sorted_numbers[third_quartile_index]

# ╔═╡ 8e17bbe4-8b39-4223-9a60-63df6d8fcfdf
# see how the values has been trimmed in a way, replaced with a threasold
begin
	sum = 0
	for i in 1:length(sorted_numbers)
		sum += 
			if sorted_numbers[i] < first_quartile_value
				first_quartile_value
			elseif sorted_numbers[i] > third_quartile_value
				third_quartile_value
			else
				sorted_numbers[i]
			end
	end
	
	sum
end

# ╔═╡ fb26e483-7022-486d-a4c8-058d46e5b10f
sum / length(sorted_numbers)

# ╔═╡ 1df7e72e-3016-4eb4-bbcd-597fa8304606
sorted_numbers[end]

# ╔═╡ b7ae4aed-5123-4698-a279-0a2cc292920e


# ╔═╡ Cell order:
# ╠═12520694-4a43-4074-a959-0c6b2fcb52e0
# ╠═30fa4d88-e32b-46a9-ad01-b87095ba223e
# ╠═020f2730-586f-4375-8153-92e1a1faeb41
# ╠═f8bc41b1-25d6-4d68-a907-c24c05a993d2
# ╠═de123813-4c9c-430a-9677-f9cb7af2ac98
# ╠═f2b31330-244f-40fe-ba97-0d5bd18408a4
# ╠═8e17bbe4-8b39-4223-9a60-63df6d8fcfdf
# ╠═fb26e483-7022-486d-a4c8-058d46e5b10f
# ╠═1df7e72e-3016-4eb4-bbcd-597fa8304606
# ╠═b7ae4aed-5123-4698-a279-0a2cc292920e
