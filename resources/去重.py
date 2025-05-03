def remove_duplicates_from_file(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    # 去重并保持原顺序
    unique_lines = list(dict.fromkeys(line.strip() for line in lines))

    with open(output_file, 'w', encoding='utf-8') as f:
        for line in unique_lines:
            f.write(line + '\n')

# 示例
input_file = 'text.txt'
output_file = '2.txt'
remove_duplicates_from_file(input_file, output_file)
print(f"去重后的内容已保存到 {output_file}")
