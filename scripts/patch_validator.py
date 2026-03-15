import hashlib
import sys
import re
import binascii

TARGET_SIMILARITY_HASH = int("0x014206767FOODDEADBEEF", 16)


def get_84_bit_hash(text_block: str) -> int:
    full_hash = hashlib.sha256(text_block.encode("utf-8")).digest()
    hash_int = int.from_bytes(full_hash[:11], byteorder="big")
    return hash_int & ((1 << 84) - 1)


def validate_payload(filepath: str, expected_github_user: str) -> bool:
    try:
        with open(filepath, "r") as f:
            content = f.read()
    except FileNotFoundError:
        print(f"[-] Error: {filepath} not found.")
        return False

    author_match = re.search(r"^From:\s+([^\s<]+)", content, re.MULTILINE)
    if not author_match:
        print("[-] Error: Could not parse 'From:' header in patch.")
        return False

    parsed_author = author_match.group(1)
    if parsed_author.lower() != expected_github_user.lower():
        print(
            f"[-] Error: Patch author '{parsed_author}' does not match PR submitter '{expected_github_user}'."
        )
        return False

    try:
        block1, rest = re.split(r"\n---\n", content, maxsplit=1)
        block2, rest = re.split(r"\n@@\s", rest, maxsplit=1)
        block3, block4 = re.split(r"\n(?=[\+\-])", rest, maxsplit=1)
        block1 = block1 + "\n---\n"
        block2 = block2 + "\n@@ "
        block3 = block3 + "\n"

    except ValueError:
        print("[-] Error: Patch does not match standard unified diff structure.")
        print(
            "    Ensure it has a '---' separator, an '@@' hunk header, and code changes (+/-)."
        )
        return False

    h1 = get_84_bit_hash(block1)
    h2 = get_84_bit_hash(block2)
    h3 = get_84_bit_hash(block3)
    h4 = get_84_bit_hash(block4)

    similarity_hash = h1 ^ h2 ^ h3 ^ h4

    print(f"[*] Parsed Author: {parsed_author}")
    print(f"[*] Computed Similarity Hash: {hex(similarity_hash)}")

    if similarity_hash == TARGET_SIMILARITY_HASH:
        print("[+] SUCCESS: Hash match! Commit approved.")
        return True
    else:
        print(f"[-] FAILURE: Hash mismatch. Expected {hex(TARGET_SIMILARITY_HASH)}.")
        return False


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python patch_validator.py <path_to_payload> <github_username>")
        sys.exit(1)

    if validate_payload(sys.argv[1], sys.argv[2]):
        sys.exit(0)
    else:
        sys.exit(1)
