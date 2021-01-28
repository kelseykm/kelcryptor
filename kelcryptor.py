#!/usr/bin/env python3

#Written by kelseykm

"""Encrypts/decrypts files with AES-256-GCM encryption"""

from getpass import getpass
import os
import readline
import glob
import time
import sys
import threading
from Crypto.Cipher import AES
from Crypto.Protocol.KDF import scrypt

version = "1.6.2"
red = "\033[1;31m"
green = "\033[1;32m"
brown = "\033[1;33m"
blue = "\033[1;34m"
white_bold = "\033[1;39m"
normal = "\033[0;39m"
cycle = "|/-\\"
cycle_stop = False
time_check = False
shred = False
banner = """%s
                                                      .-""-.                  .
                                                     / .--. \                 .
                                                    / /    \ \                .
 _  __    _  ____                  _                | |    | |                .
| |/ /___| |/ ___|_ __ _   _ _ __ | |_ ___  _ __    | |.-""-.|                .
| ' // _ \ | |   | '__| | | | '_ \| __/ _ \| '__|  ///`.::::.`\               .
| . \  __/ | |___| |  | |_| | |_) | || (_) | |    ||| ::/  \:: ;              .
|_|\_\___|_|\____|_|   \__, | .__/ \__\___/|_|    ||| ::\__/:: ;              .
                       |___/|_|                    \\\\\\ '::::' //              .
                                                    `=':-..-'`                .
%s"""%(green,normal)

class Encryptor(object):
    """Main encryption class"""

    def __init__(self, password, infile_object, outfile_object):
        self.password = password
        self.infile_object = infile_object
        self.outfile_object = outfile_object
        self.chunk_size = 51200
        self.salt_size = 32
        self.key_size = 32
        self.tag_size = 16
        self.number_of_chunks = 0
        self.number_of_read_chunks = 0

    def number_chunks(self):
        file_size = os.stat(self.infile_object.name).st_size

        if file_size == 0:
            self.number_of_chunks = file_size
        elif not file_size % self.chunk_size == 0:
            self.number_of_chunks = file_size // self.chunk_size + 1
        else:
            self.number_of_chunks = file_size/self.chunk_size

    def read_chunks(self, decrypt=False):
        if decrypt:
            chunks_left = self.number_of_chunks

        if not decrypt:
            while True:
                chunk = self.infile_object.read(self.chunk_size)
                if not chunk:
                    return
                yield chunk
                self.number_of_read_chunks += 1
        else:
            while True:
                chunk = self.infile_object.read(self.chunk_size)
                if chunks_left == 1:
                    self.tag = chunk[-self.tag_size:]
                    yield chunk[:-self.tag_size]
                    self.number_of_read_chunks += 1
                    chunks_left -= 1
                    return
                if chunks_left == 2:
                    chunk2 = self.infile_object.read(self.chunk_size)
                    if not len(chunk2) >= self.tag_size:
                        x = self.tag_size - len(chunk2)
                        self.tag = chunk[-x:] + chunk2
                        yield chunk[:-x]
                        self.number_of_read_chunks += 2
                        chunks_left -= 2
                        return
                    else:
                        self.tag = chunk2[-self.tag_size:]
                        yield chunk
                        yield chunk2[:-self.tag_size]
                        self.number_of_read_chunks += 2
                        chunks_left -= 2
                        return
                else:
                    yield chunk
                    self.number_of_read_chunks += 1
                    chunks_left -= 1

    def encryption_key(self, salt=None):
        if not salt:
            self.salt = os.urandom(self.salt_size)
        else:
            self.salt = salt

        key = scrypt(self.password, self.salt, self.key_size, N=2**15, r=8, p=1)
        return key

    def encrypt(self):
        self.nonce = os.urandom(AES.block_size)
        cipher = AES.new(self.encryption_key(), AES.MODE_GCM, self.nonce)

        self.outfile_object.write(self.salt)
        self.outfile_object.write(self.nonce)

        for chunk in self.read_chunks():
            cipher_chunk = cipher.encrypt(chunk)
            self.outfile_object.write(cipher_chunk)

        self.tag = cipher.digest()
        self.outfile_object.write(self.tag)


    def decrypt(self):
        self.salt = self.infile_object.read(self.salt_size)
        self.nonce = self.infile_object.read(AES.block_size)

        cipher = AES.new(self.encryption_key(self.salt), AES.MODE_GCM, self.nonce)

        self.number_chunks()

        self.infile_object.seek(48)
        for chunk in self.read_chunks(decrypt=True):
            clear_chunk = cipher.decrypt(chunk)
            self.outfile_object.write(clear_chunk)

        cipher.verify(self.tag)

class Shredder(object):
    """
    Main file shredder class

    CAUTION: This is meant to be used on magnetic hard drives only. It will NOT
    accomplish its function when used on solid state drives
    """

    def __init__(self, file_object):
        self.file_object = file_object
        self.file_name = file_object.name
        self.file_size = os.stat(file_object.name).st_size
        self.chunk_size = 51200

    def generate_zeros(self, number_to_generate):
        return b'\x00' * number_to_generate

    def generate_ones(self, number_to_generate):
        return b'\x01' * number_to_generate

    def generate_randoms(self, number_to_generate):
        return os.urandom(number_to_generate)

    @property
    def number_of_chunks(self):
        x = self.file_size/self.chunk_size
        if type(x) is float:
            return x.__floor__() + 1
        else:
            return x

    @property
    def length_of_last_chunk(self):
        if self.file_size % self.chunk_size == 0:
            return self.chunk_size
        else:
            return self.file_size % self.chunk_size

    def shred_file(self):
        passes = self.number_of_chunks - 1
        last = self.length_of_last_chunk

        #First pass with zeros
        self.file_object.seek(0)
        for _ in range(passes):
            zeros = self.generate_zeros(self.chunk_size)
            self.file_object.write(zeros)
        zeros = self.generate_zeros(last)
        self.file_object.write(zeros)

        #Second pass with ones
        self.file_object.seek(0)
        for _ in range(passes):
            ones = self.generate_ones(self.chunk_size)
            self.file_object.write(ones)
        ones = self.generate_ones(last)
        self.file_object.write(ones)

        #Third pass with randoms
        self.file_object.seek(0)
        for _ in range(passes):
            rands = self.generate_randoms(self.chunk_size)
            self.file_object.write(rands)
        rands = self.generate_randoms(last)
        self.file_object.write(rands)

        #Fourth and last pass with zeros
        self.file_object.seek(0)
        for _ in range(passes):
            zeros = self.generate_zeros(self.chunk_size)
            self.file_object.write(zeros)
        zeros = self.generate_zeros(last)
        self.file_object.write(zeros)

    def delete_file(self):
        os.remove(self.file_name)

class Autocomplete(object):
    """Enables autocompletion of lines"""

    def autocomplete_path(self, text, state):
        line = readline.get_line_buffer().split()
        if "~" in text:
            text = os.path.expanduser("~")
        if os.path.isdir(text):
            text += "/"
        return [x for x in glob.glob(text + "*")][state]

def progress_status(e_obj, caller):
    """Shows encryption/decryption progress"""
    while True:
        if e_obj.number_of_chunks == e_obj.number_of_read_chunks or cycle_stop:
            break
        sys.stdout.write(f"{blue}[INFO]{normal} {green}[ {white_bold}{round(e_obj.number_of_read_chunks/e_obj.number_of_chunks*100, 1)}{green} % ] {caller}...{normal}")
        sys.stdout.write("\r")
        sys.stdout.flush()

def cycler(caller):
    """Prints cycling animation as user waits for shredding to be finished"""

    while not cycle_stop:
        for c in cycle:
            sys.stdout.write(f"{blue}[INFO]{normal} ")
            sys.stdout.write(f"{green}[ {white_bold}{c}{green} ]")
            sys.stdout.write(f" {caller}...{normal}")
            sys.stdout.write("\r")
            sys.stdout.flush()
            time.sleep(0.5)

def encrypt_func(password, infile, outfile):
    """Main encryption function"""

    global cycle_stop
    cycle_stop = False

    if time_check:
        start_time = time.perf_counter()

    try:
        with open(infile, "rb") as f, open(outfile, "wb") as g:
            encryption_object = Encryptor(password, f, g)

            encryption_object.number_chunks()
            thread = threading.Thread(target=progress_status, args=(encryption_object, "ENCRYPTED"))
            thread.start()

            encryption_object.encrypt()
    except:
        cycle_stop = True
        thread.join()
        print(f"{red}[ERROR]{normal} {brown}{infile}{normal} ENCRYPTION FAILED!")
        if os.path.exists(outfile):
            os.remove(outfile)
        return

    if time_check:
        stop_time = time.perf_counter()

    thread.join()
    print(f"{blue}[INFO]{normal} {brown}{infile}{normal} ENCRYPTION SUCCESSFUL!")

    if time_check:
        print(f"{blue}[INFO]{normal} FINISHED IN {brown}{round(stop_time-start_time,2)}{normal} SECONDS")

def decrypt_func(password, infile, outfile):
    """Main decryption function"""

    global cycle_stop
    cycle_stop = False

    if time_check:
        start_time = time.perf_counter()

    try:
        with open(infile, "rb") as f, open(outfile, "wb") as g:
            decryption_object = Encryptor(password, f, g)

            decryption_object.number_chunks()
            thread = threading.Thread(target=progress_status, args=(decryption_object, "DECRYPTED"))
            thread.start()

            decryption_object.decrypt()
    except: #ValueError("MAC check failed")
        cycle_stop = True
        thread.join()

        print(f"{red}[ERROR]{normal} {brown}{infile}{normal} DECRYPTION FAILED!")
        print(f"{blue}[INFO]{normal} EITHER THE PASSWORD YOU ENTERED IS WRONG OR THE ENCRYPTED FILE HAS BEEN CORRUPTED")
        if os.path.exists(outfile):
            os.remove(outfile)
        return

    if time_check:
        stop_time = time.perf_counter()

    thread.join()
    print(f"{blue}[INFO]{normal} {brown}{infile}{normal} DECRYPTION SUCCESSFUL!")

    if time_check:
        print(f"{blue}[INFO]{normal} FINISHED IN {brown}{round(stop_time-start_time,2)}{normal} SECONDS")

def shred_func(infile):
    """Main file shredding function"""

    global cycle_stop
    cycle_stop = False

    thread = threading.Thread(target=cycler, args=("SHREDDING",))
    thread.start()

    start_time = time.perf_counter()
    with open(infile, "r+b") as f:
        shred_object = Shredder(f)
        shred_object.shred_file()
        shred_object.delete_file()

    stop_time = time.perf_counter()
    cycle_stop = True
    thread.join()
    print(f"{blue}[INFO]{normal} {white_bold}{infile}{normal} SHREDDING SUCCESSFUL!")
    if time_check:
        print(f"{blue}[INFO]{normal} FINISHED IN {brown}{round(stop_time-start_time,2)}{normal} SECONDS")

def get_intention(func):
    """Wrapper for geting user's intention"""

    def wrapper():
        print(f"""{blue}[INFO]{normal} CODES:
    1 FOR ENCRYPTION
    2 FOR DECRYPTION
    q TO CANCEL""")
        intention = input(f"{green}[INPUT]{normal} ENTER CODE: ")
        while True:
            if not intention.strip().isdigit():
                if intention == "q":
                    sys.exit()
                else:
                    print(f"{red}[ERROR]{normal} {intention} IS NOT A VALID CODE")
                    intention = input(f"{green}[INPUT]{normal} INSERT CODE: ")
            elif not 0 < int(intention) <= 2:
                print(f"{red}[ERROR]{normal} {intention} IS NOT A VALID CODE")
                intention = input(f"{green}[INPUT]{normal} INSERT CODE: ")
            else:
                break
        intention_2 = "ENCRYPTED" if intention == "1" else "DECRYPTED"
        return func(intention_2)
    return wrapper

def get_password():
    """Get password from user"""

    password = getpass(prompt=f"{green}[INPUT]{normal} PLEASE ENTER PASSWORD: ")
    while True:
        if not password:
            print(f"{red}[ERROR]{normal} BLANK PASSWORD IS NOT ALLOWED")
            password = getpass(prompt=f"{green}[INPUT]{normal} PLEASE ENTER PASSWORD: ")
        else:
            break

    password2 = getpass(prompt=f"{green}[INPUT]{normal} REPEAT PASSWORD: ")
    while True:
        if not password2:
            print(f"{red}[ERROR]{normal} BLANK PASSWORD IS NOT ALLOWED")
            password2 = getpass(prompt=f"{green}[INPUT]{normal} REPEAT PASSWORD: ")
        else:
            break

    while True:
        if not password == password2:
            print(f"{red}[ERROR]{normal} PASSWORDS DO NOT MATCH!!!")
            password = getpass(prompt=f"{green}[INPUT]{normal} PLEASE ENTER PASSWORD: ")
            password2 = getpass(prompt=f"{green}[INPUT]{normal} REPEAT PASSWORD: ")
        elif not password and not password2:
            print(f"{red}[ERROR]{normal} BLANK PASSWORD IS NOT ALLOWED")
            password = getpass(prompt=f"{green}[INPUT]{normal} PLEASE ENTER PASSWORD: ")
            password2 = getpass(prompt=f"{green}[INPUT]{normal} REPEAT PASSWORD: ")
        else:
            break

    return password

def check_file(func):
    """Wrapper for checking if files are acceptable. For multiple file input only"""

    def wrapper():
        files = sys.argv[2:] if time_check else sys.argv[1:]
        error = False
        for infile in files:
            infile = os.path.abspath(infile)
            if not os.path.exists(infile):
                print(f"{red}[ERROR]{normal} {white_bold}{infile}{normal} DOES NOT EXIST")
                error = True
            elif os.path.isdir(infile):
                print(f"{red}[ERROR]{normal} {white_bold}{infile}{normal} IS A DIRECTORY")
                error = True
            elif not os.path.isfile(infile):
                print(f"{red}[ERROR]{normal} {white_bold}{infile}{normal} IS NOT A REGULAR FILE")
                error = True
        if error:
            sys.exit()
        return func()

    return wrapper

@get_intention
def single_file_input(intention):
    """Get files to be encrypted or decrypted. For single file input only"""

    global shred

    t = Autocomplete()
    readline.set_completer_delims("\t")
    readline.parse_and_bind("tab: complete")
    readline.set_completer(t.autocomplete_path)

    if intention == "ENCRYPTED":
        shredding_intention = input(f"{green}[INPUT]{normal} WOULD YOU LIKE THE INPUT FILE TO BE SHREDDED AFTER ENCRYPTION? (Y/N) ")
        while True:
            if shredding_intention.strip().upper() == "Y" or shredding_intention.strip().upper() == "N":
                break
            else:
                print(f"{red}[ERROR]{normal} INVALID RESPONSE")
                shredding_intention = input(f"{green}[INPUT]{normal} WOULD YOU LIKE THE INPUT FILE TO BE SHREDDED AFTER ENCRYPTION? (Y/N) ")
        shred = True if shredding_intention.strip().upper() == "Y" else False

    infile = input(f"{green}[INPUT]{normal} PLEASE ENTER THE PATH OF THE FILE TO BE {intention}: ")
    while True:
        if not infile:
            print(f"{red}[ERROR]{normal} FILE NAME CANNOT BE BLANK")
            infile = input(f"{green}[INPUT]{normal} PLEASE ENTER THE PATH OF THE FILE TO BE {intention}: ")
        elif not os.path.exists(os.path.abspath(infile)):
            print(f"{red}[ERROR]{normal} THAT FILE DOES NOT EXIST")
            infile = input(f"{green}[INPUT]{normal} PLEASE ENTER THE PATH OF THE FILE TO BE {intention}: ")
        elif os.path.isdir(os.path.abspath(infile)):
            print(f"{red}[ERROR]{normal} THAT IS A DIRECTORY")
            infile = input(f"{green}[INPUT]{normal} PLEASE ENTER THE PATH OF THE FILE TO BE {intention}: ")
        elif not os.path.isfile(os.path.abspath(infile)):
            print(f"{red}[ERROR]{normal} THAT FILE IS NOT A REGULAR FILE")
            infile = input(f"{green}[INPUT]{normal} PLEASE ENTER THE PATH OF THE FILE TO BE {intention}: ")
        else:
            break

    infile = os.path.abspath(infile)
    suffix = "enc" if intention == "ENCRYPTED" else "dec"

    pre_outfile = input(f"{green}[INPUT]{normal} WHAT WOULD YOU LIKE THE {intention} FILE TO BE NAMED (LEAVE BLANK FOR DEFAULT): ")
    if pre_outfile:
        file_path, file_name = os.path.split(infile)
        outfile = f"{file_path}/{pre_outfile}.{suffix}"
    else:
        print(f"{blue}[INFO]{normal} DEFAULT NAME WILL BE USED FOR {intention} FILE")

        file_name, file_ext = os.path.splitext(infile)
        if intention == "DECRYPTED" and file_ext == ".enc":
            outfile = f"{file_name}.{suffix}"
        else:
            outfile = f"{infile}.{suffix}"

    password = get_password()
    func = encrypt_func if intention == "ENCRYPTED" else decrypt_func
    func(password, infile, outfile)

    if shred:
        shred_func(infile)

@check_file
@get_intention
def multiple_file_input(intention):
    """Get files for encryption/decryption. For multiple file input only"""

    global shred

    if intention == "ENCRYPTED":
        shredding_intention = input(f"{green}[INPUT]{normal} WOULD YOU LIKE THE INPUT FILES TO BE SHREDDED AFTER ENCRYPTION? (Y/N) ")
        while True:
            if shredding_intention.strip().upper() == "Y" or shredding_intention.strip().upper() == "N":
                break
            else:
                print(f"{red}[ERROR]{normal} INVALID RESPONSE")
                shredding_intention = input(f"{green}[INPUT]{normal} WOULD YOU LIKE THE INPUT FILES TO BE SHREDDED AFTER ENCRYPTION? (Y/N) ")
        shred = True if shredding_intention.strip().upper() == "Y" else False

    print(f"{blue}[INFO]{normal} DEFAULT NAMING WILL BE USED FOR THE {intention} FILES")
    print(f"{blue}[INFO]{normal} THE {intention} FILES WILL BE SAVED AT THEIR CURRENT LOCATION")

    password = get_password()
    suffix = "enc" if intention == "ENCRYPTED" else "dec"
    func = encrypt_func if intention == "ENCRYPTED" else decrypt_func
    files = sys.argv[2:] if time_check else sys.argv[1:]

    for infile in files:
        infile = os.path.abspath(infile)

        file_name, file_ext = os.path.splitext(infile)
        if intention == "DECRYPTED" and file_ext == ".enc":
            outfile = f"{file_name}.{suffix}"
        else:
            outfile = f"{infile}.{suffix}"

        func(password, infile, outfile)

        if shred:
            shred_func(infile)

def usage():
    """Help function. Describes how to use the application"""

    instructions = """
Usage: kelcryptor.py [OPTION] [FILES]...

kelcryptor.py encrypts and decrypts files with AES-256-GCM.

Options:

-h, --help                Show usage
-t, --time                Show time taken after encryption or decryption is finished
-v, --version             Show kelcryptor version number

Note:
    1) Input files MUST be regular files
    2) The file paths can be either relative or absolute
    3) For multi-file mode, the multiple files are added as arguments after the program name (kelcryptor.py)
    4) There MUST be more than one file for multi-file mode
    5) For single-file mode, no file arguments should be added after the program name

Examples:
  kelcryptor.py ./foo ./bar     # Start kelcryptor in multi-file mode for files foo and bar
  kelcryptor.py -h              # Show usage
  kelcryptor.py                 # Start kelcryptor in single-file mode
  kelcryptor -t ./foo ./bar     # Show time taken for multi-file mode (multi-file mode rules still apply)
  kelcryptor -t                 # Show time taken for single-file mode
"""
    print(instructions)
    sys.exit()

def main():
    """Main function"""

    global time_check

    if sys.argv[1:]:
        if len(sys.argv[1:]) == 1:
            if sys.argv[1] == "--help" or sys.argv[1] == "-h":
                usage()
            elif sys.argv[1] == "--time" or sys.argv[1] == "-t":
                time_check = True
                print(banner)
                single_file_input()
            elif sys.argv[1] == "--version" or sys.argv[1] == "-v":
                print(banner)
                print(f"{blue}[INFO]{normal} VERSION: {white_bold}{version}{normal}")
                sys.exit()
            else:
                print(f"{red}[ERROR]{normal} UNKNOWN OPTION: {sys.argv[1]}")
                usage()
        elif len(sys.argv[1:]) > 1:
            if sys.argv[1] == "--help" or sys.argv[1] == "-h":
                usage()
            elif sys.argv[1] == "--version" or sys.argv[1] == "-v":
                print(f"{red}[ERROR]{normal} VERSION CHECK OPTION DOES NOT TAKE ARGUMENTS")
                usage()
            elif sys.argv[1] == "--time" or sys.argv[1] == "-t":
                if not len(sys.argv[1:]) > 2:
                    print(f"{red}[ERROR]{normal} WRONG USAGE OF '--time' OPTION FOR MULTI-FILE MODE")
                    usage()
                else:
                    time_check = True
                    print(banner)
                    multiple_file_input()
            else:
                print(banner)
                multiple_file_input()
    else:
        print(banner)
        single_file_input()

if __name__ == "__main__":
    main()
