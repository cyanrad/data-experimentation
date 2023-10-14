numbers = [70, 72, 74, 68, 69, 87, 82, 73, 67, 70]

function winsorizedmean(numbers) -> float64
    first_quartile_index = Int(ceil(length(sorted_numbers) / 4))
    third_quartile_index = Int(ceil(length(sorted_numbers)*3 / 4))

    first_quartile_value = sorted_numbers[first_quartile_index]
    third_quartile_value = sorted_numbers[third_quartile_index]

    sorted_numbers = sort(numbers)

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

    return sum / length(sorted_numbers)
end
