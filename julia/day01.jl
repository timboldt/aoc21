module Day01

function parse_input()
    path = joinpath(@__DIR__, "..", "golang", "day01", "input.txt")
    input = open(path, "r") do file
        read(file, String)
    end
    sonar_reports = parse.(Int, split(input))
    return sonar_reports
end

function part1()
    ssrs = parse_input()
    vals = zip(ssrs[1:end-1], ssrs[2:end])
    return count(v -> (v[1] < v[2]), vals)
end

function part2()
    ssrs = parse_input()
    triples = collect(zip(ssrs[1:end-2], ssrs[2:end-1], ssrs[3:end]))
    vals = zip(triples[1:end-1], triples[2:end])
    return count(v -> (v[1][1] + v[1][2] + v[1][3] < v[2][1] + v[2][2] + v[2][3]), vals)
end

end # module

println(Day01.part1())
println(Day01.part2())